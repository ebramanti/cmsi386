#include <iostream>
using namespace std;

class nextOdd {
	private:
		int x = -1;
	public:
		int nextOddGenerator() {
			return x += 2;
		}
};

int main() {
	cout << nextOdd.nextOddGenerator();
	return 0;
}