class Solution {
public:
    int lengthLongestPath(string input) {
        int maxi=0,count=0,ln=1;
        bool isFile=false;
        vector<int> level(200);
        level[0]=0;
        for(int i=0,fin=input.size();i<fin;++i){
            //find which level
            while(input[i]=='\t'){
                ++ln;++i;
            }
            //read file name
            while(input[i]!='\n'&&i<fin){
                if(input[i]=='.')isFile=true;
                ++count;++i;
            }
            //calculate
            if(isFile){
                maxi=max(maxi,level[ln-1]+count);
            }
            else{
                level[ln]=level[ln-1]+count+1;// 1 means '/'
            }
            //reset
            count=0;ln=1;isFile=false;
        }
        return maxi;
    }
};