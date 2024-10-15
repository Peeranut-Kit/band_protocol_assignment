#include <iostream>
#include <string>
using namespace std;

bool bossBabyRevenge(const string& S) {
    // Empty string meaning no action still falls in Good boy
    if (S.empty()) return true;

    // Check if the first action is a revenge or last action is a shot
    if (S[0] == 'R' || S.back() == 'S') return false;

    int shot = 0;
    int revenge = 0;
    char prev = '\0';

    for (char ch : S) {
        if (ch == 'R') {
            revenge++;      // see R -> add revenge.
        }
        else if (ch == 'S') {
            if (prev == ch) { 
                shot++;     // sequence of S -> add shot.
            }
            else {
                // new S after previous R
                // -> If there are more S than R before this position, save remaining shot and look for next R to avenge it.
                // -> every next S cant be avenged even there are remaining R after subtracting previous S.
                shot = max(shot - revenge, 0);
                revenge = 0;
                shot++;
            }
        }
        prev = ch;
    }

    // Check at the end if all shots are avenged after last revenge
    return (revenge >= shot);
}

int main(int argc, char const *argv[])
{
    string sequence;
    cin >> sequence;

    string result = bossBabyRevenge(sequence)? "Good boy": "Bad boy";
    cout << result << endl;
    
    return 0;
}
