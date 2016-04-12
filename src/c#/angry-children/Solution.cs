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

        int[] answers = new int[n - (k - 1)];
        int j = 0;
        for (j = 0; j < n; j++)
        {
            if (j > (n - k))
                break;

            answers[j] = inputs[j + (k - 1)] - inputs[j];

            if (answers[j] == 0)
                break;
        }

        int fairness = 0;
        for (int p = 0; p < j; p++)
        {
            if (p == 0)
            {
                fairness = answers[p];
                continue;
            }

            if (answers[p] < fairness)
                fairness = answers[p];

            if (fairness == 0)
                break;
        }

        Console.WriteLine(fairness);
    }
}