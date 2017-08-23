import std.stdio;
import std.conv;

void main(string[] args) {
    int n;
    readf!"%d"(n);
    bool[] list = Eratosthenes(n);

    for(int i=0; i<n; i++)
        if(list[i]) 
            writeln(i);
}


bool[] Eratosthenes(int n){
    bool[] bList = new bool[n];
    bList[0] = false;
    bList[1] = false;

    for(int i=2; i<n; i++) 
        bList[i] = true;

    for(int i=2; i<n; i++)
        if(bList[i])
            for(int j=i*2; j<n; j+=i)
                bList[j] = false;

    return bList;
}

