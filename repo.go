package gormrepo

// Since BaseRepo is considered to have a more explicit meaning, this version renames Repo to BaseRepo.
// However, Repo is thought to be more convenient to use, so it was decided to use an alias.
// Therefore, the following code was intended to define an alias (for Go 1.23), but due to issues in GoLand, it was not possible.
// 由于认为 BaseRepo 是更具备明确意义的命名，因此这个版本把 Repo 重命名为 BaseRepo
// 但认为 Repo 这个命名更短而且也更便于使用，因此决定使用别名，这样就会很完美
// 因此使用以下代码定义别名 (go1.23)

/*
type Repo[MOD any, CLS any] = BaseRepo[MOD, CLS]

func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return NewBaseRepo[MOD, CLS]((*MOD)(nil), cls)
}
*/

// While, the above code cannot be correctly interpreted in GoLand 2025.2 EAP
// due to the existing issue https://youtrack.jetbrains.com/issue/GO-18487.
// Therefore, I had to use the following compromise version.
// 但遗憾的是前面的代码在我最喜欢的最主流编辑器 GoLand 2025.2 EAP 的版本里不能被正确解释
// 依然存在 https://youtrack.jetbrains.com/issue/GO-18487 的问题
// 因此不得不改为以下的折中版本的，这样我在其他项目里用的时候依然很方便

type Repo[MOD any, CLS any] struct {
	*BaseRepo[MOD, CLS]
}

func NewRepo[MOD any, CLS any](_ *MOD, cls CLS) *Repo[MOD, CLS] {
	return &Repo[MOD, CLS]{
		BaseRepo: NewBaseRepo[MOD, CLS]((*MOD)(nil), cls),
	}
}
