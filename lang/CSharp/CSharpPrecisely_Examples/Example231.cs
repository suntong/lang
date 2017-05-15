// Example 231 from page 191 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// RegExp -> NFA -> DFA -> Graph in Generic C#
// 2001-10-23, 2003-09-03, 2011-08-04

// This file contains, in order:
//   * A class Nfa for representing an NFA (a nondeterministic finite 
//     automaton), and for converting it to a DFA (a deterministic 
//     finite automaton).  Most complexity is in this class.
//   * A class Dfa for representing a DFA, a deterministic finite 
//     automaton, and for writing a dot input file representing the DFA.
//   * Classes for representing regular expressions, and for building an 
//     NFA from a regular expression
//   * A test class that creates an NFA, a DFA, and a dot input file 
//     for a number of small regular expressions.  The DFAs are 
//     not minimized.

using System;
using System.Text;
using SC = System.Collections;
using System.Collections.Generic;
using System.IO;

// A set, with element-based hash codes, built upon HashSet<T>

class Set<T> : IEquatable<Set<T>>, ICollection<T> where T : IEquatable<T> {
  private readonly HashSet<T> inner = new HashSet<T>();
  private int? cachedHash = null;   // Cached hash code is valid if non-null

  public Set() { }

  public Set(T x) : this() { 
    Add(x);
  }

  public Set(IEnumerable<T> coll) : this() {
    foreach (T x in coll) 
      Add(x);
  }

  public bool Contains(T x) {
    return inner.Contains(x);
  }

  public void Add(T x) {
    if (!Contains(x)) {
      inner.Add(x);
      cachedHash = null;
    }
  }

  public bool Remove(T x) {
    bool removed = inner.Remove(x);
    if (removed)
      cachedHash = null;
    return removed;
  }

  public IEnumerator<T> GetEnumerator() {
    return inner.GetEnumerator();
  }

  SC.IEnumerator SC.IEnumerable.GetEnumerator() {
    return GetEnumerator();
  }

  public int Count {
    get { return inner.Count; }
  }
  
  public void CopyTo(T[] arr, int i) {
    inner.CopyTo(arr, i);
  }

  public void Clear() {
    inner.Clear();
    cachedHash = null;
  }

  public bool IsReadOnly {
    get { return false; }
  }

  // Is this set a subset of that?
  public bool IsSubsetOf(Set<T> that) { 
    foreach (T x in this)
      if (!that.Contains(x))
        return false;
    return true;            
  }

  // Create new set as intersection of this and that
  public Set<T> Intersection(Set<T> that) { 
    Set<T> res = new Set<T>();
    foreach (T x in this)
      if (that.Contains(x))
        res.Add(x);
    return res;
  }

  // Create new set as union of this and that
  public Set<T> Union(Set<T> that) { 
    Set<T> res = new Set<T>(this);
    foreach (T x in that)
      res.Add(x);
    return res;
  }

  // Create new set as difference between this and that
  public Set<T> Difference(Set<T> that) { 
    Set<T> res = new Set<T>();
    foreach (T x in this)
      if (!that.Contains(x))
        res.Add(x);
    return res;
  }

  // Create new set as symmetric difference between this and that
  public Set<T> SymmetricDifference(Set<T> that) { 
    Set<T> res = new Set<T>();
    foreach (T x in this)
      if (!that.Contains(x))
        res.Add(x);
    foreach (T x in that)
      if (!this.Contains(x))
        res.Add(x);
    return res;
  }

  // Compute hash code based on set contents, and cache it
  public override int GetHashCode() { 
    if (!cachedHash.HasValue) {
      int res = 0;
      foreach (T x in this)
        res ^= x.GetHashCode();
      cachedHash = res;
    }
    return cachedHash.Value;
  }

  public bool Equals(Set<T> that) { 
    return that != null && that.Count == this.Count && that.IsSubsetOf(this);
  }

  public override String ToString() {
    StringBuilder res = new StringBuilder();
    res.Append("{ ");
    bool first = true;
    foreach (T x in this) {
      if (!first) 
        res.Append(", ");
      res.Append(x);
      first = false;
    }
    res.Append(" }");
    return res.ToString();
  }
}

// ----------------------------------------------------------------------

// Regular expressions, NFAs, DFAs, and dot graphs
// sestoft@itu.dk 
// Java 2001-07-10 * C# 2001-10-22 * Generic C# 2001-10-23, 2003-09-03

// We 
//  use Queue<int> and Queue<Set<int>> for worklists
//  use Set<int> for pre-DFA states
//  use List<Transition> for NFA transition relations
//  use Dictionary<Set<int>, Dictionary<String, Set<int>>>
//  use Dictionary<int, Dictionary<String, int>> for DFA transition relations
//  and we need to use Set<int> because it has a proper element-based 
//      GetHashCode(); the .NET 4.0 HashSet<T> and SortedSet<T> do not.

/* Class Nfa and conversion from NFA to DFA ---------------------------

  A nondeterministic finite automaton (NFA) is represented as a
  Map from state number (int) to a List of Transitions, a
  Transition being a pair of a label lab (a String, null meaning
  epsilon) and a target state (an int).

  A DFA is created from an NFA in two steps:

    (1) Construct a DFA whose each of whose states is composite,
        namely a set of NFA states (Set of int).  This is done by
        methods CompositeDfaTrans and EpsilonClose.

    (2) Replace composite states (Set of int) by simple states
        (int).  This is done by methods Rename and MkRenamer.

  Method CompositeDfaTrans works as follows: 

    Create the epsilon-closure S0 (a Set of ints) of the start
    state s0, and put it in a worklist (a Queue).  Create an
    empty DFA transition relation, which is a Map from a
    composite state (an epsilon-closed Set of ints) to a Map
    from a label (a non-null String) to a composite state.

    Repeatedly choose a composite state S from the worklist.  If it is
    not already in the keyset of the DFA transition relation, compute
    for every non-epsilon label lab the set T of states reachable by
    that label from some state s in S.  Compute the epsilon-closure
    Tclose of every such state T and put it on the worklist.  Then add
    the transition S -lab-> Tclose to the DFA transition relation, for
    every lab.

  Method EpsilonClose works as follows: 

    Given a set S of states.  Put the states of S in a worklist.
    Repeatedly choose a state s from the worklist, and consider all
    epsilon-transitions s -eps-> s' from s.  If s' is in S already,
    then do nothing; otherwise add s' to S and the worklist.  When the
    worklist is empty, S is epsilon-closed; return S.

  Method MkRenamer works as follows: 

    Given a Map from Set of int to something, create an
    injective Map from Set of int to int, by choosing a fresh
    int for every value of the map.

  Method Rename works as follows:

    Given a Map from Set of int to Map from String to Set of
    int, use the result of MkRenamer to replace all Sets of ints
    by ints.

*/


class Nfa {
  private readonly int startState;
  private readonly int exitState;    // This is the unique accept state
  private readonly IDictionary<int,List<Transition>> trans;

  public Nfa(int startState, int exitState) {
    this.startState = startState; this.exitState = exitState;
    trans = new Dictionary<int,List<Transition>>();
    if (!startState.Equals(exitState))
      trans.Add(exitState, new List<Transition>());
  }

  public int Start { get { return startState; } }

  public int Exit { get { return exitState; } }

  public IDictionary<int, List<Transition>> Trans { 
    get { return trans; }
  }

  public void AddTrans(int s1, String lab, int s2) {
    List<Transition> s1Trans;
    if (trans.ContainsKey(s1)) 
      s1Trans = trans[s1];
    else {
      s1Trans = new List<Transition>();
      trans.Add(s1, s1Trans);
    }
    s1Trans.Add(new Transition(lab, s2));
  }

  public void AddTrans(KeyValuePair<int, List<Transition>> tr) {
    // Assumption: if tr is in trans, it maps to an empty list (end state)
    trans.Remove(tr.Key);
    trans.Add(tr.Key, tr.Value);
  }

  public override String ToString() {
    return "NFA start=" + startState + " exit=" + exitState;
  }

  // Construct the transition relation of a composite-state DFA
  // from an NFA with start state s0 and transition relation
  // trans (a Map from int to List of Transition).  The start
  // state of the constructed DFA is the epsilon closure of s0,
  // and its transition relation is a Map from a composite state
  // (a Set of ints) to a Map from label (a String) to a
  // composite state (a Set of ints).

  static IDictionary<Set<int>, IDictionary<String, Set<int>>>
    CompositeDfaTrans(int s0, IDictionary<int, List<Transition>> trans) {
    Set<int> S0 = EpsilonClose(new Set<int>(s0), trans);
    Queue<Set<int>> worklist = new Queue<Set<int>>();
    worklist.Enqueue(S0);
    // The transition relation of the DFA
    IDictionary<Set<int>, IDictionary<String, Set<int>>> res = 
      new Dictionary<Set<int>, IDictionary<String, Set<int>>>(); 
    while (worklist.Count != 0) {
      Set<int> S = worklist.Dequeue();
      if (!res.ContainsKey(S)) {
        // The S -lab-> T transition relation being constructed for a given S
        IDictionary<String, Set<int>> STrans = 
          new Dictionary<String, Set<int>>(); 
        // For all s in S, consider all transitions s -lab-> t
        foreach (int s in S) {
          // For all non-epsilon transitions s -lab-> t, add t to T
          foreach (Transition tr in trans[s]) {
            if (tr.lab != null) {       // Already a transition on lab
              Set<int> toState; 
              if (STrans.ContainsKey(tr.lab)) 
                toState = STrans[tr.lab];
              else {                    // No transitions on lab yet
                toState = new Set<int>();
                STrans.Add(tr.lab, toState);
              }
              toState.Add(tr.target);
            }
          }
        }
        // Epsilon-close all T such that S -lab-> T, and put on worklist
        Dictionary<String, Set<int>> STransClosed = 
          new Dictionary<String, Set<int> >();
        foreach (KeyValuePair<String, Set<int>> entry in STrans) {
          Set<int> Tclose = EpsilonClose(entry.Value, trans);
          STransClosed.Add(entry.Key, Tclose);
          worklist.Enqueue(Tclose);
        }
        res.Add(S, STransClosed);
      }
    }
    return res;
  }  

  // Compute epsilon-closure of state set S in transition relation trans.  

  static Set<int> 
    EpsilonClose(Set<int> S, IDictionary<int, List<Transition>> trans) {
    // The worklist initially contains all S members
    Queue<int> worklist = new Queue<int>(S);
    Set<int> res = new Set<int>(S);
    while (worklist.Count != 0) {
      int s = worklist.Dequeue();
      foreach (Transition tr in trans[s]) {
        if (tr.lab == null && !res.Contains(tr.target)) {
          res.Add(tr.target);
          worklist.Enqueue(tr.target);
        }
      }
    }
    return res;
  }

  // Compute a renamer, which is a Map from Set of int to int

  static IDictionary<Set<int>, int> MkRenamer(ICollection<Set<int>> states) {
    IDictionary<Set<int>, int> renamer = new Dictionary<Set<int>, int>();
    int count = 0;
    foreach (Set<int> k in states) 
      renamer.Add(k, count++);
    return renamer;
  }

  // Using a renamer (a Map from Set of int to int), replace
  // composite (Set of int) states with simple (int) states in
  // the transition relation trans, which is assumed to be a Map
  // from Set of int to Map from String to Set of int.  The
  // result is a Map from int to Map from String to int.

  static IDictionary<int, IDictionary<String, int>>
    Rename(IDictionary<Set<int>, int> renamer, 
           IDictionary<Set<int>, IDictionary<String, Set<int>>> trans) {
    IDictionary<int, IDictionary<String, int>> newtrans = 
      new Dictionary<int, IDictionary<String, int>>();
    foreach (KeyValuePair<Set<int>, IDictionary<String, Set<int>>> entry 
             in trans) {
      Set<int> k = entry.Key;
      IDictionary<String, int> newktrans = new Dictionary<String, int>();
      foreach (KeyValuePair<String, Set<int>> tr in entry.Value) 
        newktrans.Add(tr.Key, renamer[tr.Value]);
      newtrans.Add(renamer[k], newktrans);
    }
    return newtrans;
  }

  static Set<int> AcceptStates(ICollection<Set<int>> states, 
                               IDictionary<Set<int>, int> renamer,
                               int exit) {
    Set<int> acceptStates = new Set<int>();
    foreach (Set<int> state in states) 
      if (state.Contains(exit)) 
        acceptStates.Add(renamer[state]);
    return acceptStates;
  }

  public Dfa ToDfa() {
    IDictionary<Set<int>, IDictionary<String, Set<int>>> 
      cDfaTrans = CompositeDfaTrans(startState, trans);
    Set<int> cDfaStart = EpsilonClose(new Set<int>(startState), trans);
    ICollection<Set<int>> cDfaStates = cDfaTrans.Keys;
    IDictionary<Set<int>, int> renamer = MkRenamer(cDfaStates);
    IDictionary<int, IDictionary<String, int>> simpleDfaTrans = 
      Rename(renamer, cDfaTrans);
    int simpleDfaStart = renamer[cDfaStart];
    Set<int> simpleDfaAccept = AcceptStates(cDfaStates, renamer, exitState);
    return new Dfa(simpleDfaStart, simpleDfaAccept, simpleDfaTrans);
  }

  // Nested class for creating distinctly named states when constructing NFAs

  public class NameSource {
    private static int nextName = 0;

    public int next() { 
      return nextName++; 
    }
  }
}

// Class Transition, a transition from one state to another ----------

  public class Transition {
    public String lab;
    public int target;
    
    public Transition(String lab, int target) { 
      this.lab = lab; this.target = target; 
    }
    
    public override String ToString() {
      return "-" + lab + "-> " + target;
    }
  }

// Class Dfa, deterministic finite automata --------------------------

/*
  A deterministic finite automaton (DFA) is represented as a Map
  from state number (int) to a Map from label (a String,
  non-null) to a target state (an int).  
*/

class Dfa {
  private readonly int startState;
  private readonly Set<int> acceptStates;
  private readonly IDictionary<int, IDictionary<String,int>> trans;

  public Dfa(int startState, Set<int> acceptStates, 
             IDictionary<int, IDictionary<String,int>> trans) {
    this.startState = startState; 
    this.acceptStates = acceptStates;
    this.trans = trans;
  }
  
  public int Start { get { return startState; } }

  public Set<int> Accept { get { return acceptStates; } }

  public IDictionary<int, IDictionary<String,int>> Trans { 
    get { return trans; }
  }

  public override String ToString() {
    return "DFA start=" + startState + "\naccept=" + acceptStates;
  }

  // Write an input file for the dot program.  You can find dot at
  // http://www.research.att.com/sw/tools/graphviz/

  public void WriteDot(String filename) {
    TextWriter wr = 
      new StreamWriter(new FileStream(filename, FileMode.Create, 
                                      FileAccess.Write));
    wr.WriteLine("// Format this file as a Postscript file with ");
    wr.WriteLine("//    dot " + filename + " -Tps -o out.ps\n");
    wr.WriteLine("digraph dfa {");
    wr.WriteLine("size=\"11,8.25\";");
    wr.WriteLine("rotate=90;");
    wr.WriteLine("rankdir=LR;");
    wr.WriteLine("n999999 [style=invis];");    // Invisible start node
    wr.WriteLine("n999999 -> n" + startState); // Edge into start state
    
    // Accept states are double circles
    foreach (int state in trans.Keys) 
      if (acceptStates.Contains(state))
        wr.WriteLine("n" + state + " [peripheries=2];");

    // The transitions 
    foreach (KeyValuePair<int, IDictionary<String, int>> entry in trans) {
      int s1 = entry.Key;
      foreach (KeyValuePair<String, int> s1Trans in entry.Value) {
        String lab = s1Trans.Key;
        int s2 = s1Trans.Value;
        wr.WriteLine("n" + s1 + " -> n" + s2 + " [label=\"" + lab + "\"];");
      }
    }
    wr.WriteLine("}");
    wr.Close();
  }
}

// Regular expressions ----------------------------------------------
//
// Abstract syntax of regular expressions
//    r ::= A | r1 r2 | (r1|r2) | r*
//

abstract class Regex { 
  abstract public Nfa MkNfa(Nfa.NameSource names);
}

class Eps : Regex {
  // The resulting nfa0 has form s0s -eps-> s0e

  public override Nfa MkNfa(Nfa.NameSource names) {
    int s0s = names.next();
    int s0e = names.next();
    Nfa nfa0 = new Nfa(s0s, s0e);
    nfa0.AddTrans(s0s, null, s0e);
    return nfa0;
  }
}

class Sym : Regex {
  String sym;

  public Sym(String sym) { 
    this.sym = sym; 
  }

  // The resulting nfa0 has form s0s -sym-> s0e

  public override Nfa MkNfa(Nfa.NameSource names) {
    int s0s = names.next();
    int s0e = names.next();
    Nfa nfa0 = new Nfa(s0s, s0e);
    nfa0.AddTrans(s0s, sym, s0e);
    return nfa0;
  }
}

class Seq : Regex {
  Regex r1, r2;

  public Seq(Regex r1, Regex r2) {
    this.r1 = r1; this.r2 = r2;
  }

  // If   nfa1 has form s1s ----> s1e 
  // and  nfa2 has form s2s ----> s2e 
  // then nfa0 has form s1s ----> s1e -eps-> s2s ----> s2e

  public override Nfa MkNfa(Nfa.NameSource names) {
    Nfa nfa1 = r1.MkNfa(names);
    Nfa nfa2 = r2.MkNfa(names);
    Nfa nfa0 = new Nfa(nfa1.Start, nfa2.Exit);
    foreach (KeyValuePair<int, List<Transition>> entry in nfa1.Trans) 
      nfa0.AddTrans(entry);
    foreach (KeyValuePair<int, List<Transition>> entry in nfa2.Trans) 
      nfa0.AddTrans(entry);
    nfa0.AddTrans(nfa1.Exit, null, nfa2.Start);
    return nfa0;
  }
}

class Alt : Regex {
  Regex r1, r2;

  public Alt(Regex r1, Regex r2) {
    this.r1 = r1; this.r2 = r2;
  }

  // If   nfa1 has form s1s ----> s1e 
  // and  nfa2 has form s2s ----> s2e 
  // then nfa0 has form s0s -eps-> s1s ----> s1e -eps-> s0e
  //                    s0s -eps-> s2s ----> s2e -eps-> s0e

  public override Nfa MkNfa(Nfa.NameSource names) {
    Nfa nfa1 = r1.MkNfa(names);
    Nfa nfa2 = r2.MkNfa(names);
    int s0s = names.next();
    int s0e = names.next();
    Nfa nfa0 = new Nfa(s0s, s0e);
    foreach (KeyValuePair<int, List<Transition>> entry in nfa1.Trans) 
      nfa0.AddTrans(entry);
    foreach (KeyValuePair<int, List<Transition>> entry in nfa2.Trans) 
      nfa0.AddTrans(entry);
    nfa0.AddTrans(s0s, null, nfa1.Start);
    nfa0.AddTrans(s0s, null, nfa2.Start);
    nfa0.AddTrans(nfa1.Exit, null, s0e);
    nfa0.AddTrans(nfa2.Exit, null, s0e);
    return nfa0;
  }
}

class Star : Regex {
  Regex r;

  public Star(Regex r) {
    this.r = r; 
  }

  // If   nfa1 has form s1s ----> s1e 
  // then nfa0 has form s0s ----> s0s
  //                    s0s -eps-> s1s
  //                    s1e -eps-> s0s

  public override Nfa MkNfa(Nfa.NameSource names) {
    Nfa nfa1 = r.MkNfa(names);
    int s0s = names.next();
    Nfa nfa0 = new Nfa(s0s, s0s);
    foreach (KeyValuePair<int, List<Transition>> entry in nfa1.Trans) 
      nfa0.AddTrans(entry);
    nfa0.AddTrans(s0s, null, nfa1.Start);
    nfa0.AddTrans(nfa1.Exit, null, s0s);
    return nfa0;
  }
}

// Trying the RE->NFA->DFA translation on three regular expressions

class TestNFA {
  public static void Main(String[] args) {
    Regex a = new Sym("A");
    Regex b = new Sym("B");
    Regex c = new Sym("C");
    Regex abStar = new Star(new Alt(a, b));
    Regex bb = new Seq(b, b);
    Regex r = new Seq(abStar, new Seq(a, b));
    // The regular expression (a|b)*ab
    BuildAndShow("dfa1.dot", r);
    // The regular expression ((a|b)*ab)*
    BuildAndShow("dfa2.dot", new Star(r));
    // The regular expression ((a|b)*ab)((a|b)*ab)
    BuildAndShow("dfa3.dot", new Seq(r, r));
    // The regular expression (a|b)*abb, from ASU 1986 p 136
    BuildAndShow("dfa4.dot", new Seq(abStar, new Seq(a, bb)));
    // SML reals: sign?((digit+(\.digit+)?))([eE]sign?digit+)?
    Regex d = new Sym("digit");
    Regex dPlus = new Seq(d, new Star(d));
    Regex s = new Sym("sign");
    Regex sOpt = new Alt(s, new Eps());
    Regex dot = new Sym(".");
    Regex dotDigOpt = new Alt(new Eps(), new Seq(dot, dPlus));
    Regex mant = new Seq(sOpt, new Seq(dPlus, dotDigOpt));
    Regex e = new Sym("e");
    Regex exp = new Alt(new Eps(), new Seq(e, new Seq(sOpt, dPlus)));
    Regex smlReal = new Seq(mant, exp);
    BuildAndShow("dfa5.dot", smlReal);
  }

  public static void BuildAndShow(String filename, Regex r) {
    Nfa nfa = r.MkNfa(new Nfa.NameSource());
    Console.WriteLine(nfa);
    Console.WriteLine("---");
    Dfa dfa = nfa.ToDfa();
    Console.WriteLine(dfa);
    Console.WriteLine("Writing DFA graph to file " + filename);
    dfa.WriteDot(filename);
    Console.WriteLine();
  }
}
