package main

type ASTNode struct {
   token    Token
   level    int
   children []ASTNode
}

func (n *ASTNode) add(t Token){
   new_node := ASTNode{t, n.level, []ASTNode{}}
   n.children = append(n.children, new_node)
}

func (n *ASTNode) addNode(new_node ASTNode){
   n.children = append(n.children, new_node)
}

func (n ASTNode) String() string{
   s String
   for i := 0; i < n.level; i++ {
      s += "    "
   }
}
