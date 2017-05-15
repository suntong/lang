// Example 105 from page 87 of C# Precisely, 2nd ed. (MIT Press 2012)
// Authors: Peter Sestoft (sestoft@itu.dk) and Henrik I. Hansen

using System;
using System.Collections.Generic;
using System.Text;

public class SparseMatrix {
  private readonly int rows;
  // A SparseMatrix has an array of lists of NonZeros, one for each column.
  // Invariant: In each list the nonzeros appear in increasing order of nz.i
  readonly List<NonZero>[] cols;
  
  // Create a sparse matrix from 2D array B which must be rectangular

  public SparseMatrix(double[][] B) {
    cols = new List<NonZero>[B.Length];
    rows = B.Length != 0 ? B[0].Length : 0;
    for (int j=0; j<B.Length; j++) {
      cols[j] = new List<NonZero>();
      for (int i=0; i<rows; i++) 
        if (B[i][j] != 0.0)
          cols[j].Add(new NonZero(i, B[i][j]));
    }
  }

  // Create an all-zero rows-by-cols sparse matrix

  public SparseMatrix(int r, int c) {
    cols = new List<NonZero>[c];
    this.rows = r;
    for (int j=0; j<c; j++) 
      cols[j] = new List<NonZero>();
  }
    
  // Properties to get the number of rows and columns of the matrix

  public int Rows {
    get { return rows; }
  }

  public int Cols {
    get { return cols.Length; }
  }

  // Indexer to get and set an element of the matrix

  public double this[int i, int j] {
    get {
      List<NonZero> colj = this[j];
      int k = 0; 
      while (k < colj.Count && colj[k].i < i)
	k++;
      return k < colj.Count && colj[k].i == i ? colj[k].Mij : 0.0;
    }
    set {
      List<NonZero> colj = this[j];
      int k = 0; 
      while (k < colj.Count && colj[k].i < i)
	k++;
      if (k < colj.Count && colj[k].i == i) 
        colj[k].Mij = value;
      else if (value != 0.0)
        colj.Insert(k, new NonZero(i, value));
    }
  }
  
  // Indexer to get j'th column of matrix
  
  private List<NonZero> this[int j] {
    get { return cols[j]; }
  }

  // A pair of a row number i and a non-zero element B[i][-]

  private class NonZero {
    public readonly int i;
    public double Mij;

    public NonZero(int i, double Mij) {
      this.i = i; this.Mij = Mij;
    }

    public NonZero(NonZero nz) {
      this.i = nz.i; this.Mij = nz.Mij;
    }
  }

  public static SparseMatrix Add(SparseMatrix A, SparseMatrix B) {
    if (A.Rows == B.Rows && A.Cols == B.Cols) {
      int rRows = A.Rows, rCols = A.Cols;
      SparseMatrix R = new SparseMatrix(rRows, rCols);
      for (int j=0; j<rCols; j++) {
        List<NonZero> Aj = A[j], Bj = B[j], Rj = R[j];
        int ak = 0, bk = 0;
        while (ak<Aj.Count && bk<Bj.Count) {
          if (Aj[ak].i < Bj[bk].i) 
            Rj.Add(new NonZero(Aj[ak++]));
          else if (Bj[bk].i < Aj[ak].i)
            Rj.Add(new NonZero(Bj[bk++]));
          else // Aj[ak].i==Bj[bk].i
            Rj.Add(new NonZero(Aj[ak].i, Aj[ak++].Mij+Bj[bk++].Mij));
        } 
        while (ak<Aj.Count) 
          Rj.Add(new NonZero(Aj[ak++]));
        while (bk<Bj.Count) 
          Rj.Add(new NonZero(Bj[bk++]));
      }
      return R;
    } else
      throw new ApplicationException("SparseMatrix.Add: Matrix size misfit");
  }

  public override String ToString() {
    StringBuilder sb = new StringBuilder();
    for (int i=0; i<Rows; i++) {
      for (int j=0; j<Cols; j++)
        sb.AppendFormat("{0,6} ", this[i,j]);
      sb.Append("\n");
    }
    return sb.ToString();
  }
}

class TestSparseMatrix {
  public static void Main(String[] args) {
    SparseMatrix A = new SparseMatrix(4, 5), B = new SparseMatrix(4, 5);
    A[0,2] = 102; A[0,3] = 103; A[1,0] = 110; A[3,4] = 134;
    B[0,2] = 202; B[1,3] = 213; B[2,0] = 220; B[3,4] = 234;
    Console.WriteLine("A =\n{0}", A);
    Console.WriteLine("B =\n{0}", B);
    Console.WriteLine("A+B =\n{0}", SparseMatrix.Add(A,B));
  }
}
