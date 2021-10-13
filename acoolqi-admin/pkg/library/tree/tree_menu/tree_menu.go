package tree_menu

import (
	"acoolqi-admin/models"
	"github.com/yuanhao2015/acoolTools"
	"sort"
)

// Tree 统一定义菜单树的数据结构，也可以自定义添加其他字段
type Tree struct {
	Name            string      `json:"name,omitempty"` //节点名字
	Path            string      `json:"path,omitempty"`
	Hidden          bool        `json:"hidden"`
	Redirect        string      `json:"redirect,omitempty"`
	Component       string      `json:"component,omitempty"`
	AlwaysShow      bool        `json:"alwaysShow"`
	Data            interface{} `json:"-"` //自定义对象
	MenuType        string      `json:"menuType"`
	Meta            interface{} `json:"meta"`
	Leaf            bool        `json:"-"`                  //叶子节点
	Selected        bool        `json:"-"`                  //选中
	PartialSelected bool        `json:"-"`                  //部分选中
	Children        []Tree      `json:"children,omitempty"` //子节点
	Id              int         `json:"id"`
	Label           string      `json:"label"`
}

// ConvertToINodeArray 其他的结构体想要生成菜单树，直接实现这个接口
type INode interface {
	// GetName 获取显示名字
	GetName() string
	// GetMenuId 获取id
	GetMenuId() int
	// GetParentId 获取父id
	GetParentId() int
	// GetData 获取附加数据
	GetData() interface{}
	// IsRoot 判断当前节点是否是顶层根节点
	IsRoot() bool
	// GetPath 路径
	GetPath() string
	// GetMenuType
	GetMenuType() string
	GetId() int
	GetLabel() string
}
type INodes []INode

func (nodes INodes) Len() int {
	return len(nodes)
}
func (nodes INodes) Swap(i, j int) {
	nodes[i], nodes[j] = nodes[j], nodes[i]
}
func (nodes INodes) Less(i, j int) bool {
	return nodes[i].GetMenuId() < nodes[j].GetMenuId()
}

// GenerateTree 自定义的结构体实现 INode 接口后调用此方法生成树结构
// nodes 需要生成树的节点
// selectedNode 生成树后选中的节点
// menuTrees 生成成功后的树结构对象
func GenerateTree(nodes, selectedNodes []INode) (trees []Tree) {
	trees = []Tree{}
	// 定义顶层根和子节点
	var roots, childs []INode
	for _, v := range nodes {
		if v.IsRoot() {
			// 判断顶层根节点
			roots = append(roots, v)
		}
		childs = append(childs, v)
	}

	for _, v := range roots {
		var flag = false

		var component = "Layout"

		m := make(map[string]interface{})
		if v.GetData() != nil {
			menu := v.GetData().(models.SysMenu)
			m["title"] = menu.MenuName
			m["icon"] = menu.Icon
			m["noCache"] = menu.IsCache == 1
			if menu.IsFrame == 0 {
				m["link"] = menu.Path
			} else {
				m["link"] = nil
			}
			flag = menu.Visible == "1"
			if !acoolTools.StrUtils.HasEmpty(menu.Component) {
				component = menu.Component
			}
		}
		childTree := &Tree{
			Name:       v.GetName(),
			Data:       v.GetData(),
			Path:       "/" + v.GetPath(),
			Hidden:     flag,
			AlwaysShow: true,
			Redirect:   "noRedirect",
			Meta:       m,
			Component:  component,
			Id:         v.GetId(),
			Label:      v.GetLabel(),
		}
		// 递归之前，根据父节点确认 childTree 的选中状态
		childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children)
		// 递归
		recursiveTree(childTree, childs, selectedNodes)
		// 递归之后，根据子节点确认 childTree 的选中状态
		if !childTree.Selected {
			childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children)
		}
		// 递归之后，根据子节点确认 childTree 的半选中状态
		childTree.PartialSelected = nodePartialSelected(childTree.Children)
		// 递归之后，根据子节确认是否是叶子节点
		childTree.Leaf = len(childTree.Children) == 0
		trees = append(trees, *childTree)
	}
	return
}

// recursiveTree 递归生成树结构
// tree 递归的树对象
// nodes 递归的节点
// selectedNodes 选中的节点
func recursiveTree(tree *Tree, nodes, selectedNodes []INode) {
	data := tree.Data.(INode)

	for _, v := range nodes {
		if v.IsRoot() {
			// 如果当前节点是顶层根节点就跳过
			continue
		}
		var flag = false
		var component string
		var redirect string
		if v.GetParentId() == 1 && v.GetMenuType() == "M" {
			component = "ParentView"
			redirect = "noRedirect"
		}
		m := make(map[string]interface{})
		if v.GetData() != nil {
			menu := v.GetData().(models.SysMenu)
			flag = menu.Visible == "1"
			m["title"] = menu.MenuName
			m["icon"] = menu.Icon
			m["noCache"] = menu.IsCache == 1
			if menu.IsFrame == 0 {
				m["link"] = menu.Path
			} else {
				m["link"] = nil
			}
			if !acoolTools.StrUtils.HasEmpty(menu.Component) {
				component = menu.Component
			}
		}
		if data.GetMenuId() == v.GetParentId() {
			childTree := &Tree{
				Name:      v.GetName(),
				Data:      v.GetData(),
				Path:      v.GetPath(),
				Hidden:    flag,
				Meta:      m,
				Redirect:  redirect,
				Component: component,
				MenuType:  v.GetMenuType(),
				Id:        v.GetId(),
				Label:     v.GetLabel(),
			}
			// 递归之前，根据子节点和父节点确认 childTree 的选中状态
			childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children) || tree.Selected
			recursiveTree(childTree, nodes, selectedNodes)

			if !childTree.Selected {
				// 递归之后，根据子节点确认 childTree 的选中状态
				childTree.Selected = nodeSelected(v, selectedNodes, childTree.Children)
			}
			// 递归之后，根据子节点确认 childTree 的半选中状态
			childTree.PartialSelected = nodePartialSelected(childTree.Children)
			// 递归之后，根据子节确认是否是叶子节点
			childTree.Leaf = len(childTree.Children) == 0
			tree.Children = append(tree.Children, *childTree)
		}
	}
}

// FindRelationNode 在 allTree 中查询 nodes 中节点的所有父节点
// nodes 要查询父节点的子节点数组
// allTree 所有节点数组
func FindRelationNode(nodes, allNodes []INode) (respNodes []INode) {
	nodeMap := make(map[int]INode)
	for _, v := range nodes {
		recursiveFindRelationNode(nodeMap, allNodes, v, 0)
	}

	for _, v := range nodeMap {
		respNodes = append(respNodes, v)
	}
	sort.Sort(INodes(respNodes))
	return
}

// recursiveFindRelationNode 递归查询关联父子节点
// nodeMap 查询结果搜集到map中
// allNodes 所有节点
// node 递归节点
// t 递归查找类型：0 查找父、子节点；1 只查找父节点；2 只查找子节点
func recursiveFindRelationNode(nodeMap map[int]INode, allNodes []INode, node INode, t int) {
	nodeMap[node.GetMenuId()] = node
	for _, v := range allNodes {
		if _, ok := nodeMap[v.GetMenuId()]; ok {
			continue
		}
		// 查找父节点
		if t == 0 || t == 1 {
			if node.GetParentId() == v.GetMenuId() {
				nodeMap[v.GetMenuId()] = v
				if v.IsRoot() {
					// 是顶层根节点时，不再进行递归
					continue
				}
				recursiveFindRelationNode(nodeMap, allNodes, v, 1)
			}
		}
		// 查找子节点
		if t == 0 || t == 2 {
			if node.GetMenuId() == v.GetParentId() {
				nodeMap[v.GetMenuId()] = v
				recursiveFindRelationNode(nodeMap, allNodes, v, 2)
			}
		}
	}
}

// nodeSelected 判断节点的选中状态
// node 进行判断节点
func nodeSelected(node INode, selectedNodes []INode, children []Tree) bool {
	for _, v := range selectedNodes {
		if node.GetMenuId() == v.GetMenuId() {
			// 1. 如果选择节点数组中存在当前节点
			return true
		}
	}

	if len(children) == 0 {
		// 2. 不满足前置条件1，且没有子节点
		return false
	}
	selectedNum := 0
	for _, v := range children {
		if v.Selected {
			selectedNum++
		}
	}
	if selectedNum == len(children) {
		// 不满足前置条件1，2 ，且子节点全部是选中状态
		return true
	}
	return false
}

// nodePartialSelected 判断节点的半选中状态
func nodePartialSelected(trees []Tree) bool {
	selectedNum := 0
	for _, v := range trees {
		if v.Selected {
			selectedNum++
		}
	}
	if selectedNum == len(trees) || selectedNum == 0 {
		// 子节点全选中，或一个也没有选中
		return false
	}
	return true
}

type SystemMenus []models.SysMenu

// ConvertToINodeArray 将当前数组转换成父类 INode 接口 数组
func (s SystemMenus) ConvertToINodeArray(*[]models.SysMenu) (nodes []INode) {
	for _, v := range s {
		nodes = append(nodes, v)
	}
	return
}

// GetTree 获取树结构数据
func (s SystemMenus) GetTree(m *[]models.SysMenu) []Tree {
	s = *m
	array := s.ConvertToINodeArray(m)
	return GenerateTree(array, nil)
}
