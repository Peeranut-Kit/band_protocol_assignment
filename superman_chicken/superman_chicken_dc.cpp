#include <iostream>
#include <vector>
using namespace std;

int saved_chicken_number(int k, vector<int> &chicken_pos, int start, int stop) {
    // base case
    if (start == stop) return 1;

    // mid-point
    int m = (start+stop)/2;

    // find maximum chicken number saved from left and right halves
    int left = saved_chicken_number(k, chicken_pos, start, m);
    int right = saved_chicken_number(k, chicken_pos, m+1, stop);

    // find maximum chicken number saved in case we place the roof cover the mid-point
    int max_mid = 1;
    for (int i = m; i >= start; i--)
    {
        int j = m+1;
        if (chicken_pos[i] + k - 1 < chicken_pos[j]) continue;

        while (j < stop && chicken_pos[i] + k - 1 >= chicken_pos[j+1]) j++;
        max_mid = max(max_mid, j - i + 1);
    }

    return max(max(left,right),max_mid);
}

int main(int argc, char const *argv[])
{
    int n, k;
    cin >> n >> k;

    vector<int> chicken_pos(n);
    for (size_t i = 0; i < n; i++)
    {
        cin >> chicken_pos[i];
    }

    cout << saved_chicken_number(k, chicken_pos, 0, n-1) << endl;
    return 0;
}
