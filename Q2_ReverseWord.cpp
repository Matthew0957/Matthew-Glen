#include <iostream>
#include <string>

using namespace std;

const int MAX_SIZE = 100;
char charStack[MAX_SIZE];
int top = -1;

void push(char c)
{
    if (top < MAX_SIZE - 1)
    {
        charStack[++top] = c;
    }
}

char pop()
{
    if (top >= 0)
    {
        return charStack[top--];
    }
    return '\0';
}

int main()
{
    string word;
    cout << "Enter a word: ";
    cin >> word;


    for (int i = 0; i < word.length(); i++)
    {
        push(word[i]);
    }

    cout << "Original word: " << word << "\n";
    cout << "Reversed word: ";


    for (int i = 0; i < word.length(); i++)
    {
        cout << pop();
    }
    cout << "\n";

    return 0;
}
