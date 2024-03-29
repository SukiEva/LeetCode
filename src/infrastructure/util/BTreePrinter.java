package infrastructure.util;

import infrastructure.pojo.TreeNode;

import java.util.ArrayList;
import java.util.Collections;
import java.util.List;

/**
 * 二叉树 树形打印
 * <p>
 * <a href=
 * "https://stackoverflow.com/questions/4965335/how-to-print-binary-tree-diagram-in-java">StackOverFlow</a>
 */
class BTreePrinter {
    public static void print(TreeNode root) {
        int maxLevel = BTreePrinter.maxLevel(root);
        printNodeInternal(Collections.singletonList(root), 1, maxLevel);
    }

    private static void printNodeInternal(List<TreeNode> nodes, int level, int maxLevel) {
        if (nodes.isEmpty() || BTreePrinter.isAllElementsNull(nodes)) {
            return;
        }
        int floor = maxLevel - level;
        int endgeLines = (int) Math.pow(2, (Math.max(floor - 1, 0)));
        int firstSpaces = (int) Math.pow(2, (floor)) - 1;
        int betweenSpaces = (int) Math.pow(2, (floor + 1)) - 1;
        BTreePrinter.printWhitespaces(firstSpaces);
        List<TreeNode> newNodes = new ArrayList<>();
        for (TreeNode node : nodes) {
            if (node != null) {
                System.out.print(node.val);
                newNodes.add(node.left);
                newNodes.add(node.right);
            } else {
                newNodes.add(null);
                newNodes.add(null);
                System.out.print(" ");
            }
            BTreePrinter.printWhitespaces(betweenSpaces);
        }
        System.out.println();
        for (int i = 1; i <= endgeLines; i++) {
            for (TreeNode node : nodes) {
                BTreePrinter.printWhitespaces(firstSpaces - i);
                if (node == null) {
                    BTreePrinter.printWhitespaces(endgeLines + endgeLines + i + 1);
                    continue;
                }
                if (node.left != null) {
                    System.out.print("/");
                } else {
                    BTreePrinter.printWhitespaces(1);
                }
                BTreePrinter.printWhitespaces(i + i - 1);
                if (node.right != null) {
                    System.out.print("\\");
                } else {
                    BTreePrinter.printWhitespaces(1);
                }
                BTreePrinter.printWhitespaces(endgeLines + endgeLines - i);
            }
            System.out.println();
        }
        printNodeInternal(newNodes, level + 1, maxLevel);
    }

    private static void printWhitespaces(int count) {
        if (count > 0) {
            System.out.print(" ".repeat(count));
        }
    }

    private static int maxLevel(TreeNode node) {
        if (node == null) {
            return 0;
        }
        return Math.max(BTreePrinter.maxLevel(node.left), BTreePrinter.maxLevel(node.right)) + 1;
    }

    private static boolean isAllElementsNull(List<TreeNode> list) {
        for (Object object : list) {
            if (object != null) {
                return false;
            }
        }
        return true;
    }
}