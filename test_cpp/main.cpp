# include <iostream>
#include <chrono>
#include <thread>

using namespace std;
int main() {
  string s;
  while(cin >> s) {
  	std::this_thread::sleep_for(std::chrono::milliseconds(3000));
    cout << s.size() << endl;
  }
}