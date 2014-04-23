using System;

namespace Snail2
{
    class Program
    {
        static void Main(string[] args)
        {
            int x = 0;

            if (args.Length == 0)
            {
                Console.WriteLine("Please input a inter which is less than 20;");
                x = Convert.ToInt16(Console.ReadLine());
                Console.Write(x);
                Console.WriteLine();
            }
            else
            {
                x = Convert.ToInt16(args[0]);
            }

            int[,] Ma = new int[20, 20];
            Matrix(Ma, 0, x);

            for (int i = 0; i < x; i++)
            {
                for (int j = 0; j < x; j++)
                {
                    Console.Write("{0,4}", Ma[i, j]);
                }
                Console.WriteLine();
            }
            Console.Read();
        }

        static void Matrix(int[,] Ma, int start, int x)
        {
            if (x == 0)
            {
            }
            else if (x == 1)
            {
                Ma[start, start] = 1;
            }
            else
            {
                for (int i = 0; i < x - 1; i++)
                {
                    Ma[start + i, start] = x * x - i;
                    Ma[start + x - 1, start + i] = x * x - (x - 1) - i;
                    Ma[start + x - 1 - i, start + x - 1] = x * x - 2 * (x - 1) - i;
                    Ma[start, start + x - 1 - i] = x * x - 3 * (x - 1) - i;
                }
                Matrix(Ma, start + 1, x - 2);
            }
        }
    }
}
