#include <iostream>
#include <vector>
using namespace std;

int saved_chicken_number(int k, vector<int> &chicken_pos, int start, int n) {
    int max_chicken = 0;
    for (int trav = 0; trav < n; trav++) {
        while (start < n && chicken_pos[start] + k <= chicken_pos[trav]) start++;
        max_chicken = max(max_chicken, trav - start + 1);
    }
    return max_chicken;
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

    cout << saved_chicken_number(k, chicken_pos, 0, n) << endl;
    return 0;
}
