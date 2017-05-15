// Example 253 from page 215 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

// Compile with 
//   
//    csc /d:DEBUG Example253.cs

using System;
using SC = System.Collections;          // IEnumerator
using System.Collections.Generic;       // IEnumerable<T>, Queue<T>
using System.Diagnostics;               // Debug
using System.IO;                        // TextReader, TextWriter
using System.Text;                      // StringBuilder
using System.Text.RegularExpressions;   // Regex

class Example253 {
  public static void Main(String[] args) {
    if (args.Length != 2) 
      Console.WriteLine("Usage: Example253 <textfile> <linewidth>\n");
    else {
      IEnumerator<String> words = 
        new ReaderEnumerator(new StreamReader(args[0]));
      int lineWidth = int.Parse(args[1]);
      Format(words, lineWidth, Console.Out);
    }
  }

  // This method formats a sequence of words (Strings obtained from
  // an enumerator) into lines of text, padding the inter-word
  // gaps with extra spaces to obtain lines of length lineWidth, and 
  // thus a straight right margin.
  //
  // There are the following exceptions: 
  //  * if a word is longer than lineWidth, it is put on a line by
  //    itself (producing a line longer than lineWidth)
  //  * a single word may appear on a line by itself (producing a line
  //     shorter than lineWidth) if adding the next word to the line
  //     would make the line longer than lineWidth
  //  * the last line of the output is not padded with extra spaces.
  //
  // The algorithm for padding with extra spaces ensures that the
  // spaces are evenly distributed over inter-word gaps, using modulo
  // arithmetics.  An Assert method call asserts that the resulting
  // output line has the correct length unless the line contains only
  // a single word or is the last line of the output.

  public static void Format(IEnumerator<String> words, int lineWidth, 
                            TextWriter tw) {
    lineWidth = Math.Max(0, lineWidth);
    WordList curLine = new WordList();
    bool moreWords = words.MoveNext();
    while (moreWords) {
      while (moreWords && curLine.Length < lineWidth) {
        String word = words.Current;
        if (word != null && word != "") 
          curLine.AddLast(word);
        moreWords = words.MoveNext();
      }
      int wordCount = curLine.Count;
      if (wordCount > 0) {
        int extraSpaces = lineWidth - curLine.Length;
        if (wordCount > 1 && extraSpaces < 0) { // last word goes on next line
          int lastWordLength = curLine.GetLast.Length;
          extraSpaces += 1 + lastWordLength;
          wordCount -= 1;
        } else if (!moreWords)                  // last line, do not pad
          extraSpaces = 0;
        // Pad inter-word space with evenly distributed extra blanks
        int holes = wordCount - 1;
        int spaces = holes/2;
        StringBuilder sb = new StringBuilder();
        sb.Append(curLine.RemoveFirst());
        for (int i=1; i<wordCount; i++) {
          spaces += extraSpaces;
          appendBlanks(sb, 1 + spaces / holes);
          spaces %= holes;
          sb.Append(curLine.RemoveFirst());
        }
        String res = sb.ToString();
        Debug.Assert(res.Length==lineWidth || wordCount==1 || !moreWords);
        tw.WriteLine(res);
      }
    }
    tw.Flush();
  }

  private static void appendBlanks(StringBuilder sb, int count) {
    for (int i=0; i<count; i++)
      sb.Append(' ');
  }
}

// A word list with a fast length method, and invariant assertions

class WordList {
  private Queue<String> strings = new Queue<String>();
  // Invariant: length equals word lengths plus inter-word spaces
  private int length = -1;   
  private String lastAdded = null;

  public int Length { get { return length; } }

  public int Count { get { return strings.Count; } }
  
  public void AddLast(String s) {
    lastAdded = s;
    strings.Enqueue(s);
    length += 1 + s.Length;
    Debug.Assert(length == computeLength() + strings.Count - 1);
  }

  public String RemoveFirst() {
    String res = strings.Dequeue();
    length -= 1 + res.Length;
    Debug.Assert(length == computeLength() + strings.Count - 1);
    return res;
  }

  public String GetLast {
    get { return lastAdded; }
  }

  private int computeLength() {  // For checking the invariant only
    int sum = 0;
    foreach (String s in strings) 
      sum += s.Length;
    return sum;
  }
}

// A String-producing IEnumerator, created from a TextReader

class ReaderEnumerator : IEnumerator<String> {
  private static Regex delim = new Regex("[ \\t]+");
  private TextReader rd;
  private String[] thisLine = null;
  private int available = 0;

  public ReaderEnumerator(TextReader rd) {
    this.rd = rd;
  }

  public bool MoveNext() {
    available--;
    // If necessary, try to find some non-blank words
    String line;
    while (rd != null && available <= 0 && null != (line = rd.ReadLine())) {
      thisLine = delim.Split(line);
      available = thisLine.Length;
    }
    return available >= 1;
  }
        
  public String Current {
    get { 
      if (available >= 1) 
        return thisLine[thisLine.Length-available];
      else 
        throw new InvalidOperationException();
    }
  }

  public void Dispose() {
    if (rd != null) {
      rd.Close();
      available = 0;
    }
    rd = null;
  }

  Object SC.IEnumerator.Current {
    get { return Current; }
  }

  void SC.IEnumerator.Reset() {
    throw new NotSupportedException();
  }
}
