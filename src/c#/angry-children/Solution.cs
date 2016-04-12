// https://www.hackerrank.com/challenges/angry-children
using System;
class Solution
{
    static void Main(string[] args)
    {
        int n = Convert.ToInt32(Console.ReadLine());
        int k = Convert.ToInt32(Console.ReadLine());
        int[] inputs = new int[n];

        for (int i = 0; i < n; i++)
        {
            int b = Convert.ToInt32(Console.ReadLine());
            inputs[i] = b;
        }

        Array.Sort(inputs);
        int fairness = inputs[n - 1] - inputs[0];
        for (int j = 0; j < n; j++)
        {
            if (j > (n - k))
                break;

            int f = inputs[j + (k - 1)] - inputs[j];
            if (f < fairness) {
                fairness = f;
            }                

            if (fairness == 0)
                break;
        }
        Console.WriteLine(fairness);
    }
}