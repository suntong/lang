// Example 194 from page 161 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.IO;        // Directory, DirectoryInfo, FileInfo

public class DirectoryHierarchyExample {
  public static void ShowDir(int indent, DirectoryInfo dir) {
    Indent(indent); Console.WriteLine(dir.Name);
    DirectoryInfo[] subdirs = dir.GetDirectories();
    foreach (DirectoryInfo d in subdirs)
      ShowDir(indent+4, d);
    FileInfo[] files = dir.GetFiles();
    foreach (FileInfo file in files) {
      Indent(indent); Console.WriteLine(file.Name);
    }
  }

  public static void Indent(int indent) {
    for (int i=0; i<indent; i++)
      Console.Write('-');
  }

  public static void Main() {
    DirectoryInfo dir = new DirectoryInfo(Directory.GetCurrentDirectory());
    ShowDir(0, dir);
  }
}
