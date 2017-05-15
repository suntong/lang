// Example 23 from page 21 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;

public class StringEks {
  public static void Main() {
    String text =
      @"C# is a class-based single-inheritance object-oriented programming
language designed for the Common Language 
Runtime of Microsoft's .Net
platform, a managed execution environment with a
typesafe intermediate language and automatic memory management.  Thus
C# is similar to the Java programming language in many respects, but
it is different in almost all details.  In general, C# favors
programmer convenience over language simplicity.  It was designed by
Anders Hejlsberg, Scott Wiltamuth and Peter Golde from Microsoft
Corporation.";
    Console.WriteLine(Readability(text));
  }
        
  static double Readability(String text) {
    int wordCount = 0, longWordsCount = 0;
    String[] sentences = text.Split(new char[] {'.'});             // Split into sentences
    foreach (String sentence in sentences) {
      String[] words = sentence.Split(' ', ',');                   // Split into words
      // String[] words = sentence.Split(new char[] {' ', ','});   // Alternative
      wordCount += words.Length;
      foreach (String word in words) {
        if (word.Length > 6)
          longWordsCount++;
        else if (word.Length == 0)
          wordCount--;                                          
      }
    }
    return (wordCount*1.0)/sentences.Length + (longWordsCount*100.0)/wordCount;
  }
}
