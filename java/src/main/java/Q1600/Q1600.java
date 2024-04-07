package Q1600;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

public class Q1600 {
}

class ThroneInheritance {

    private class Node {
        String name;
        List<Node> children = new ArrayList<>();
    }

    private Map<String, Node> name2Node = new HashMap<>();
    private Node root = new Node();
    private Map<String, Boolean> death = new HashMap<>();

    public ThroneInheritance(String kingName) {
        name2Node.put(kingName, root);
        root.name = kingName;
    }

    public void birth(String parentName, String childName) {
        Node p = name2Node.get(parentName);
        Node newNode = new Node();
        newNode.name = childName;
        name2Node.put(childName, newNode);
        p.children.add(newNode);
    }

    public void death(String name) {
        death.put(name, true);
    }

    public List<String> getInheritanceOrder() {
        List<String> res = new ArrayList<>();
        successor(root, res);
        return res;
    }

    private void successor(Node node, List<String> order) {
        if (node == null) return;
        if (!death.containsKey(node.name)) {
            order.add(node.name);
        }
        for (Node child : node.children) {
            successor(child, order);
        }
    }
}
