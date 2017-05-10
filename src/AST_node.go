package main

type ASTNode struct {
   token    Token
   level    int
   children []ASTNode
}

func Node() {
   n := ASTNode{}
   n.token = Token{"NONE","",0,0}
}

func TokenNode(token Token) ASTNode {
   n := ASTNode{}
   n.token = token
   return n
}

func (n *ASTNode) add(t Token){
   new_node := ASTNode{t, n.level, []ASTNode{}}
   n.children = append(n.children, new_node)
}

func (n *ASTNode) addNode(new_node ASTNode){
   n.children = append(n.children, new_node)
}

func (n ASTNode) String() string{
   s := ""
   s += n.token.String()
   for _,e := range n.children {
      s += e.String()
   }
   return s

}
