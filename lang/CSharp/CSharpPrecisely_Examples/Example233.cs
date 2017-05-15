// Example 233 from page 193 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Representing and traversing a directed graph with unlabelled edges.

using System;                           // Console
using System.Collections.Generic;       // Dictionary, Queue, Stack

// Graph nodes labelled with a T value

public class Node<T> {
  private readonly T label;
  private Node<T>[] neighbors; 
  
  public Node(T label) : this(label, new Node<T>[] { }) { }
  
  public Node(T label, Node<T>[] neighbors) {
    this.label = label; 
    this.neighbors = neighbors;
  }
  
  public Node<T>[] Neighbors {
    get { return neighbors; }
    set { neighbors = value; }
  }

  // Visit all nodes reachable from root.
  // The Dictionary is used as a set of nodes; only the key matters,  
  // and the value false associated with every key is ignored.
  // Using Queue (and Enqueue, Dequeue) gives breadth-first traversal
  // Using Stack (and Push, Pop) gives depth-first traversal

  public void VisitBreadthFirst() {
    Dictionary<Node<T>,bool> visited = new Dictionary<Node<T>,bool>();
    Queue<Node<T>> worklist = new Queue<Node<T>>();
    visited.Add(this, false);
    worklist.Enqueue(this);
    // Invariant: every node in the worklist is also in the visited set
    while (worklist.Count != 0) {
      Node<T> node = worklist.Dequeue();
      Console.Write("{0} ", node.label);
      foreach (Node<T> neighbor in node.Neighbors) 
        if (!visited.ContainsKey(neighbor)) {
          visited.Add(neighbor, false);
          worklist.Enqueue(neighbor);
        }
    }
    Console.WriteLine();
  }

  public void VisitDepthFirst() {
    Dictionary<Node<T>,bool> visited = new Dictionary<Node<T>,bool>();
    Stack<Node<T>> worklist = new Stack<Node<T>>();
    visited.Add(this, false);
    worklist.Push(this);
    // Invariant: every node in the worklist is also in the visited set
    while (worklist.Count != 0) {
      Node<T> node = worklist.Pop();
      Console.Write("{0} ", node.label);
      foreach (Node<T> neighbor in node.Neighbors) 
        if (!visited.ContainsKey(neighbor)) {
          visited.Add(neighbor, false);
          worklist.Push(neighbor);
        }
    }
    Console.WriteLine();
  }
}

class TestGraphs {
  public static void Main(String[] args) {
    Node<int> 
      leaf4 = new Node<int>(4), leaf5 = new Node<int>(5), 
      leaf6 = new Node<int>(6), leaf7 = new Node<int>(7);
    Node<int> 
      node2 = new Node<int>(2, new Node<int>[] { leaf4, leaf5 }),
      node3 = new Node<int>(3, new Node<int>[] { leaf6, leaf7 });
    Node<int> tree = new Node<int>(1, new Node<int>[] { node2, node3 });
    tree.VisitBreadthFirst();                           // 1 2 3 4 5 6 7
    tree.VisitDepthFirst();                             // 1 3 7 6 2 5 4
    Node<String> 
      v1 = new Node<String>("one"),
      v2 = new Node<String>("two"),
      v3 = new Node<String>("three"),
      v4 = new Node<String>("four");
    v1.Neighbors = new Node<String>[] { v1, v2, v3 };
    v2.Neighbors = new Node<String>[] { v4 };
    v3.Neighbors = new Node<String>[] { v4 };
    v4.Neighbors = new Node<String>[] { v2 };
    v1.VisitBreadthFirst();                             // one two three four
    v1.VisitDepthFirst();                               // one three four two
  }
}
