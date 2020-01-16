class Solution {
public:
	vector<double> calcEquation(vector<vector<string>>& equations, vector<double>& values, vector<vector<string>>& queries) {
		vector<double> res;
		unordered_map<string, vector<pair<string, double>> > adj;
		unordered_map<string, int > vis;
		int n = equations.size();
		for (int i = 0; i < n; i++) {
			vector<string> e = equations[i];
			double val = values[i];
			string a = e[0], b = e[1];
			adj[a].push_back(make_pair(b, val));
			adj[b].push_back(make_pair(a, 1.0 / val));
			vis[a] = 0;
			vis[b] = 0;
		}

		int m = queries.size();
		for (int i = 0; i < m; i++) {
			for (auto it = vis.begin(); it != vis.end(); it++)
				it->second = 0;
			string a = queries[i][0], b = queries[i][1];
			vis[a] = 1;
			double d = query(a, b, 1.0, adj, vis);
			res.push_back(d);
		}

		return res;
	}

	double query(string start, string end, double d, unordered_map<string, vector<pair<string, double>> > &adj, unordered_map<string, int>& vis) {
		if (adj.count(start) == 0 || adj.count(end) == 0)
			return -1.0;
		if (start == end) {
			return d;
		}

		vector<pair<string, double>> a = adj[start];
		for (int i = 0; i < a.size(); i++) {
			string ns = a[i].first;
			double val = a[i].second;
			if (vis[ns] == 1)
				continue;
			vis[ns] = 1;
			double res = query(ns, end, d*val, adj, vis);
			if (res != -1.0)
				return res;
			vis[ns] = 0;
		}
		return -1.0;
	}
};