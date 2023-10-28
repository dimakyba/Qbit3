using System;
using System.Linq;
class Program {
  static void Main(string[] args)
  {
    int n = Int32.Parse(Console.ReadLine());

    int[] data = Console.ReadLine().Split(' ').Select(Int32.Parse).ToArray();


    int length = data.Length;

  for (int i = 1; i < n; i++)
  {
    for (int j = 0; i < length; j++)
    {
      Console.Write($"{data[j]} ");
    }
    Console.WriteLine();
  }
  }
}
