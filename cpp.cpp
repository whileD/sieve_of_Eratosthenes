#include <iostream>
#include <vector>

using std::vector;

vector<bool> Eratosthenes(int);

int main() {
	int n;
	std::cin >> n;
	vector<bool> list = Eratosthenes(n);

	for(int i=0; i<n; i++)
		if(list[i])
			std::cout << i << std::endl;
}

vector<bool> Eratosthenes(int n){
	vector<bool> bList(n, true);
	bList[0] = false;
	bList[1] = false;

	for(int i=0; i<n; i++)
		if(bList[i])
			for(int j=i*2; j<n; j+=i)
				bList[j] = false;

	return bList;
}
