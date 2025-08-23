package example13_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"github.com/yyle88/done"
	"github.com/yyle88/gormrepo"
	"github.com/yyle88/gormrepo/internal/examples/example13/internal/models"
	"github.com/yyle88/rese"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var testDB *gorm.DB

func TestMain(m *testing.M) {
	db := rese.P1(gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	}))
	defer rese.F0(rese.P1(db.DB()).Close)

	done.Done(db.AutoMigrate(&models.Account{}, &models.Transaction{}))

	// 创建测试账户
	accounts := []*models.Account{
		{AccountNumber: "ACC001", AccountName: "Alice Smith", Balance: 10000.00, AccountType: "CURRENT", Status: "ACTIVE", BankCode: "BANK001", BranchCode: "BR001"},
		{AccountNumber: "ACC002", AccountName: "Bob Johnson", Balance: 5000.00, AccountType: "SAVINGS", Status: "ACTIVE", BankCode: "BANK001", BranchCode: "BR002"},
		{AccountNumber: "ACC003", AccountName: "Charlie Brown", Balance: 15000.00, AccountType: "CURRENT", Status: "ACTIVE", BankCode: "BANK002", BranchCode: "BR001"},
		{AccountNumber: "ACC004", AccountName: "Diana Prince", Balance: 0.00, AccountType: "SAVINGS", Status: "ACTIVE", BankCode: "BANK001", BranchCode: "BR001"},
	}

	for _, account := range accounts {
		done.Done(db.Create(account).Error)
	}

	testDB = db
	m.Run()
}

// TestSimpleTransaction 演示简单的事务操作
func TestSimpleTransaction(t *testing.T) {
	accountRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Account{}))
	txnRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Transaction{}))

	// 开启事务
	err := testDB.Transaction(func(tx *gorm.DB) error {

		// 创建新账户
		newAccount := &models.Account{
			AccountNumber: "ACC999",
			AccountName:   "Test User",
			Balance:       1000.00,
			AccountType:   "SAVINGS",
			Status:        "ACTIVE",
			BankCode:      "BANK001",
			BranchCode:    "BR001",
		}

		err := tx.Create(newAccount).Error
		if err != nil {
			return err
		}

		// 记录交易
		transaction := &models.Transaction{
			TransactionID:   fmt.Sprintf("TXN-%d", time.Now().UnixNano()),
			ToAccountNumber: newAccount.AccountNumber,
			Amount:          1000.00,
			TransactionType: "DEPOSIT",
			Description:     "Initial deposit",
			Status:          "COMPLETED",
			Reference:       "INIT001",
		}

		return tx.Create(transaction).Error
	})

	require.NoError(t, err)

	// 验证账户已创建
	account, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq("ACC999"))
	})
	require.NoError(t, err)
	require.NotNil(t, account)
	require.Equal(t, "Test User", account.AccountName)

	// 验证交易记录已创建
	var count int64
	err = txnRepo.Where(func(db *gorm.DB, cls *models.TransactionColumns) *gorm.DB {
		return db.Where(cls.ToAccountNumber.Eq("ACC999")).Model(&models.Transaction{})
	}).Count(&count).Error
	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	t.Logf("Successfully created account %s with initial deposit", account.AccountNumber)
}

// TestMoneyTransfer 演示转账事务（最经典的事务场景）
func TestMoneyTransfer(t *testing.T) {
	accountRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Account{}))
	txnRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Transaction{}))

	fromAccount := "ACC001"
	toAccount := "ACC002"
	transferAmount := 1500.00

	// 获取转账前的余额
	fromAccountBefore, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(fromAccount))
	})
	require.NoError(t, err)

	toAccountBefore, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(toAccount))
	})
	require.NoError(t, err)

	t.Logf("Before transfer - From: $%.2f, To: $%.2f", fromAccountBefore.Balance, toAccountBefore.Balance)

	// 执行转账事务
	err = testDB.Transaction(func(tx *gorm.DB) error {
		txAccountRepo := gormrepo.NewGormRepo(gormrepo.Use(tx, &models.Account{}))

		// 1. 从源账户扣款
		err := txAccountRepo.Where(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq(fromAccount)).Where(cls.Balance.Gte(transferAmount)).Model(&models.Account{})
		}).Updates(map[string]interface{}{
			"balance": gorm.Expr("balance - ?", transferAmount),
		}).Error
		if err != nil {
			return fmt.Errorf("failed to debit from account: %w", err)
		}

		// 检查是否有记录被更新
		var updatedCount int64
		var account models.Account
		cls := account.Columns()
		tx.Model(&models.Account{}).Where(cls.AccountNumber.Eq(fromAccount)).Where(cls.Balance.Gte(0)).Count(&updatedCount)
		if updatedCount == 0 {
			return fmt.Errorf("insufficient funds in account %s", fromAccount)
		}

		// 2. 向目标账户加款
		err = txAccountRepo.Where(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq(toAccount)).Model(&models.Account{})
		}).Updates(map[string]interface{}{
			"balance": gorm.Expr("balance + ?", transferAmount),
		}).Error
		if err != nil {
			return fmt.Errorf("failed to credit to account: %w", err)
		}

		// 3. 记录转账交易
		transactionID := fmt.Sprintf("TXN-%d", time.Now().UnixNano())
		transaction := &models.Transaction{
			TransactionID:     transactionID,
			FromAccountNumber: fromAccount,
			ToAccountNumber:   toAccount,
			Amount:            transferAmount,
			TransactionType:   "TRANSFER",
			Description:       fmt.Sprintf("Transfer from %s to %s", fromAccount, toAccount),
			Status:            "COMPLETED",
			Reference:         fmt.Sprintf("REF-%d", time.Now().Unix()),
		}

		return tx.Create(transaction).Error
	})

	require.NoError(t, err)

	// 验证转账后的余额
	fromAccountAfter, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(fromAccount))
	})
	require.NoError(t, err)

	toAccountAfter, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(toAccount))
	})
	require.NoError(t, err)

	t.Logf("After transfer - From: $%.2f, To: $%.2f", fromAccountAfter.Balance, toAccountAfter.Balance)

	// 验证余额变化正确
	require.Equal(t, fromAccountBefore.Balance-transferAmount, fromAccountAfter.Balance)
	require.Equal(t, toAccountBefore.Balance+transferAmount, toAccountAfter.Balance)

	// 验证交易记录
	var count int64
	err = txnRepo.Where(func(db *gorm.DB, cls *models.TransactionColumns) *gorm.DB {
		return db.Where(cls.FromAccountNumber.Eq(fromAccount)).
			Where(cls.ToAccountNumber.Eq(toAccount)).
			Where(cls.Amount.Eq(transferAmount)).
			Where(cls.Status.Eq("COMPLETED")).Model(&models.Transaction{})
	}).Count(&count).Error
	require.NoError(t, err)
	require.Equal(t, int64(1), count)

	t.Logf("Transfer completed successfully: $%.2f from %s to %s", transferAmount, fromAccount, toAccount)
}

// TestFailedTransaction 演示事务回滚（余额不足场景）
func TestFailedTransaction(t *testing.T) {
	accountRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Account{}))

	fromAccount := "ACC002"
	toAccount := "ACC003"
	transferAmount := 10000.00 // 超过余额的金额

	// 获取转账前的余额
	fromAccountBefore, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(fromAccount))
	})
	require.NoError(t, err)

	toAccountBefore, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(toAccount))
	})
	require.NoError(t, err)

	t.Logf("Before failed transfer attempt - From: $%.2f, To: $%.2f", fromAccountBefore.Balance, toAccountBefore.Balance)

	// 尝试转账（应该失败）
	err = testDB.Transaction(func(tx *gorm.DB) error {
		txAccountRepo := gormrepo.NewGormRepo(gormrepo.Use(tx, &models.Account{}))

		// 检查余额是否足够
		account, err := txAccountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq(fromAccount))
		})
		if err != nil {
			return err
		}

		if account.Balance < transferAmount {
			return fmt.Errorf("insufficient funds: need $%.2f, have $%.2f", transferAmount, account.Balance)
		}

		// 这里不会执行到，因为余额检查会失败
		return nil
	})

	require.Error(t, err)
	require.Contains(t, err.Error(), "insufficient funds")
	t.Logf("Transaction correctly failed: %v", err)

	// 验证余额没有变化
	fromAccountAfter, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(fromAccount))
	})
	require.NoError(t, err)

	toAccountAfter, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq(toAccount))
	})
	require.NoError(t, err)

	// 余额应该保持不变
	require.Equal(t, fromAccountBefore.Balance, fromAccountAfter.Balance)
	require.Equal(t, toAccountBefore.Balance, toAccountAfter.Balance)

	t.Logf("After failed transfer attempt - From: $%.2f, To: $%.2f (unchanged)", fromAccountAfter.Balance, toAccountAfter.Balance)
}

// TestBatchOperationsInTransaction 演示事务中的批量操作
func TestBatchOperationsInTransaction(t *testing.T) {
	accountRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Account{}))
	txnRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Transaction{}))

	// 批量奖金发放
	bonusAmount := 500.00
	accounts := []string{"ACC001", "ACC002", "ACC003"}

	// 获取发放前余额
	balancesBefore := make(map[string]float64)
	for _, accNum := range accounts {
		account, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq(accNum))
		})
		require.NoError(t, err)
		balancesBefore[accNum] = account.Balance
	}

	// 批量发放奖金事务
	err := testDB.Transaction(func(tx *gorm.DB) error {
		txAccountRepo := gormrepo.NewGormRepo(gormrepo.Use(tx, &models.Account{}))

		for _, accNum := range accounts {
			// 给每个账户增加奖金
			err := txAccountRepo.Where(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
				return db.Where(cls.AccountNumber.Eq(accNum)).Model(&models.Account{})
			}).Updates(map[string]interface{}{
				"balance": gorm.Expr("balance + ?", bonusAmount),
			}).Error
			if err != nil {
				return fmt.Errorf("failed to add bonus to account %s: %w", accNum, err)
			}

			// 记录每笔奖金交易
			transaction := &models.Transaction{
				TransactionID:   fmt.Sprintf("BONUS-%s-%d", accNum, time.Now().UnixNano()),
				ToAccountNumber: accNum,
				Amount:          bonusAmount,
				TransactionType: "BONUS",
				Description:     "Monthly bonus payment",
				Status:          "COMPLETED",
				Reference:       fmt.Sprintf("BONUS-%d", time.Now().Unix()),
			}

			err = tx.Create(transaction).Error
			if err != nil {
				return fmt.Errorf("failed to record bonus transaction for %s: %w", accNum, err)
			}
		}

		return nil
	})

	require.NoError(t, err)

	// 验证所有账户都收到奖金
	for _, accNum := range accounts {
		account, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq(accNum))
		})
		require.NoError(t, err)
		require.Equal(t, balancesBefore[accNum]+bonusAmount, account.Balance)
		t.Logf("Account %s: Before $%.2f, After $%.2f", accNum, balancesBefore[accNum], account.Balance)
	}

	// 验证奖金交易记录
	var bonusCount int64
	err = txnRepo.Where(func(db *gorm.DB, cls *models.TransactionColumns) *gorm.DB {
		return db.Where(cls.TransactionType.Eq("BONUS")).Where(cls.Status.Eq("COMPLETED")).Model(&models.Transaction{})
	}).Count(&bonusCount).Error
	require.NoError(t, err)
	require.Equal(t, int64(len(accounts)), bonusCount)

	t.Logf("Successfully distributed bonus of $%.2f to %d accounts", bonusAmount, len(accounts))
}

// TestNestedTransactionBehavior 演示事务嵌套和保存点行为
func TestNestedTransactionBehavior(t *testing.T) {
	accountRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Account{}))

	// 获取初始余额
	account, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq("ACC001"))
	})
	require.NoError(t, err)
	initialBalance := account.Balance

	// 外层事务
	err = testDB.Transaction(func(tx1 *gorm.DB) error {
		txAccountRepo1 := gormrepo.NewGormRepo(gormrepo.Use(tx1, &models.Account{}))

		// 第一次操作：增加100
		err := txAccountRepo1.Where(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq("ACC001")).Model(&models.Account{})
		}).Updates(map[string]interface{}{
			"balance": gorm.Expr("balance + ?", 100.0),
		}).Error
		if err != nil {
			return err
		}

		// 内层事务（实际上是保存点）
		return tx1.Transaction(func(tx2 *gorm.DB) error {
			txAccountRepo2 := gormrepo.NewGormRepo(gormrepo.Use(tx2, &models.Account{}))

			// 第二次操作：增加200
			err := txAccountRepo2.Where(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
				return db.Where(cls.AccountNumber.Eq("ACC001")).Model(&models.Account{})
			}).Updates(map[string]interface{}{
				"balance": gorm.Expr("balance + ?", 200.0),
			}).Error
			if err != nil {
				return err
			}

			// 内层事务成功
			return nil
		})
	})

	require.NoError(t, err)

	// 验证最终余额
	finalAccount, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq("ACC001"))
	})
	require.NoError(t, err)

	expectedBalance := initialBalance + 300.0 // 100 + 200
	require.Equal(t, expectedBalance, finalAccount.Balance)

	t.Logf("Nested transaction completed - Initial: $%.2f, Final: $%.2f, Change: +$300.00",
		initialBalance, finalAccount.Balance)
}

// TestConcurrentTransactions 演示并发事务的处理
func TestConcurrentTransactions(t *testing.T) {
	// 注意：这个测试在内存SQLite中可能不会完全展示并发问题
	// 在实际的多连接数据库中会更有意义

	accountRepo := gormrepo.NewGormRepo(gormrepo.Use(testDB, &models.Account{}))

	// 获取初始余额
	account, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq("ACC004"))
	})
	require.NoError(t, err)
	initialBalance := account.Balance

	// 模拟两个并发的存款操作
	depositAmount := 100.0

	// 第一个存款事务
	err1 := testDB.Transaction(func(tx *gorm.DB) error {
		txAccountRepo := gormrepo.NewGormRepo(gormrepo.Use(tx, &models.Account{}))

		return txAccountRepo.Where(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq("ACC004")).Model(&models.Account{})
		}).Updates(map[string]interface{}{
			"balance": gorm.Expr("balance + ?", depositAmount),
		}).Error
	})

	// 第二个存款事务
	err2 := testDB.Transaction(func(tx *gorm.DB) error {
		txAccountRepo := gormrepo.NewGormRepo(gormrepo.Use(tx, &models.Account{}))

		return txAccountRepo.Where(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
			return db.Where(cls.AccountNumber.Eq("ACC004")).Model(&models.Account{})
		}).Updates(map[string]interface{}{
			"balance": gorm.Expr("balance + ?", depositAmount),
		}).Error
	})

	require.NoError(t, err1)
	require.NoError(t, err2)

	// 验证最终余额
	finalAccount, err := accountRepo.First(func(db *gorm.DB, cls *models.AccountColumns) *gorm.DB {
		return db.Where(cls.AccountNumber.Eq("ACC004"))
	})
	require.NoError(t, err)

	expectedBalance := initialBalance + 200.0 // 两次存款
	require.Equal(t, expectedBalance, finalAccount.Balance)

	t.Logf("Concurrent transactions completed - Initial: $%.2f, Final: $%.2f, Expected: $%.2f",
		initialBalance, finalAccount.Balance, expectedBalance)
}
