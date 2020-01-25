public class TreeTest
{  public static void main(String[] args)
   {  Tree staff = new Tree();
      staff.insert("Romeo");
      staff.insert("Juliet");
      staff.insert("Tom");
      staff.insert("Dick");
      staff.insert("Harry");

      staff.print();
      staff.printSmallest();
   }
}

/**
   This class implements a binary search tree whose
   nodes hold objects that implement the Comparable
   interface.
*/

class Tree
{  /**
      Constructs an empty tree.
   */
   public Tree()
   {  root = null;
   }
   
   /**
      Inserts a new node into the tree.
      @param obj the object to insert
   */
   public void insert(Comparable obj) 
   {  Node newNode = new Node();
      newNode.data = obj;
      newNode.left = null;
      newNode.right = null;
      if (root == null) root = newNode;
      else root.insertNode(newNode);
   }
   
   /**
      Prints the contents of the tree in sorted order.
   */
   public void print()
   {  if (root != null)
         root.printNodes();
   }
  
   /**
      Prints the smallest of the tree.
   */
   public void printSmallest()
  {
    System.out.print("\nThe smallest is: " + smallest());
    }
  
  /**
      Return the smallest element.
   */
   public Comparable smallest()
   {
     if (root != null)
       return root.smallest().data;
     else
       return null;
   }
  
   private Node root;

   private class Node
   {  /**
         Inserts a new node as a descendant of this node.
         @param newNode the node to insert
      */
      public void insertNode(Node newNode)
      {  if (newNode.data.compareTo(data) < 0)
         {  if (left == null) left = newNode;
            else left.insertNode(newNode);
         }
         else
         {  if (right == null) right = newNode;
            else right.insertNode(newNode);
         }
      }

      /**
         Prints this node and all of its descendants
         in sorted order.
      */
      public void printNodes()
      {
	if (left != null)
	  left.printNodes();
	printNode();
	if (right != null)
	  right.printNodes();
      }
   
       /**
         Prints this node
      */
      public void printNode()
      {
	System.out.println(data);
      }
   
  /**
      Return the smallest element.
   */
   public Node smallest()
   {
     //System.out.println("] " + data);
     if (left != null)
       return left.smallest();
     else
       return this;
   }
  
      public Comparable data;
      public Node left;
      public Node right;
   }
}
