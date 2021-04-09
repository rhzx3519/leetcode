class TreeNode(object):

    def __init__(self, name):
        self.name = name
        self.child = []
        self.dead = False

    def insert(self, parentName, childName):
        def insert(root):
            if not root:
                return
            if root.name==parentName:
                node = TreeNode(childName)
                root.child.append(node)
                return node
            for ch in root.child:
                insert(ch)
        return insert(self)


    def remove(self, name):
        def remove(root):
            if not root:
                return
            if root.name==name:
                root.dead = True
                return
            for ch in root.child:
                remove(ch)
        remove(self)

    def traverse(self):
        ans = []
        def traverse(root):
            if not root:
                return
            if not root.dead:
                ans.append(root)
            for ch in root.child:
                traverse(ch)
        traverse(self)
        return ans
    
    def present(self):
        return ' -> '.join(map(lambda node: node.name, self.traverse()))

class ThroneInheritance(object):

    def __init__(self, kingName):
        """
        :type kingName: str
        """
        self.root = TreeNode(kingName)
        self.mem = {kingName: self.root}


    def birth(self, parentName, childName):
        """
        :type parentName: str
        :type childName: str
        :rtype: None
        """
        node = self.mem[parentName].insert(parentName, childName)
        self.mem[childName] = node

    def death(self, name):
        """
        :type name: str
        :rtype: None
        """
        self.mem[name].dead = True

    def getInheritanceOrder(self):
        """
        :rtype: List[str]
        """
        return map(lambda root: root.name, self.root.traverse())



# Your ThroneInheritance object will be instantiated and called as such:
# obj = ThroneInheritance(kingName)
# obj.birth(parentName,childName)
# obj.death(name)
# param_3 = obj.getInheritanceOrder()


if __name__ == '__main__':
    root = TreeNode('King')
    root.insert('King', 'Alice')
    print root.present()
    # print root.name, root.child[0].name











