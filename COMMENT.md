# problem1
3と5の剰余で計算すればOK

# problem2
フィボナッチ数を毎回計算しつつ、値が偶数なら足す。
3n+2の項が偶数なのがわかれば手計算でもできそう

# problem3
エラトステネスの篩を実装して、値の平方根までの素数を算出した後、
順番に剰余をとって、0になるものを出していく。
愚直にやると実行が遅いので、因数判定するときは毎回平方根をとって枝刈りしたほうがいい
素因数分解を実装するのはそれなりに面倒

# problem4
全探索で計算
最大値を記録するようにして、
掛け算の値が最大値以下の場合は判定しないロジックを追加

# problem5
手計算なら簡単にできそう。
結局素因数分解を実装する話

# problem6
愚直に実装して終了。
手計算だとどうやるんだろう、二項定理がアップを始めそう

# problem7
problem3で実装したエラトステネスの篩を少し改造して完了
実行時間は思ったより早い

# problem8
文字列とみなしてソートを使ってごにょる
0が含まれていると必ず0なので、枝刈りしている

# problem9
愚直に実装。
x^2+y^2 < (x+y)^2 (x,y is Natural)を使って少しだけ計算を削減
ループの抜け方があまりよくない

# problem10
エラトステネスの篩を改造するだけ。
対象の数字の平方根までしか見ないようにしないとすごく遅い
行挿入以前は30秒以上、挿入すると1,5秒くらいで完了する

# problem11
整数問題ではなく2次元的な問題。
斜めを1通りしか検討しておらず、正解を見つけるのに手間取った。
ムダに長いコードになってしまっており反省。

# problem12
素因数分解して、約数の数を愚直に数えるアプローチ。
素数テーブルをいくつまで作るかが悩みどころだったが、
一番最初に約数が出現するという性格上、約数が低めの値に集まる可能性が高い。
手元で検証したところ、1000くらいの素数テーブルでも正しい答えが導けた。

# problem13
愚直に実装。
合計値を算出した後数値を文字列扱いして先頭10文字を切り取る。
言語によっては64bitでも足せないということもあるのだろうか…

追記：
golangで書き換えた。10桁5分割してint64で演算を行い、繰り上がりを適用した。
繰り上がりを適用しないと、末尾2桁の数字が異なる（実際は1だけ違う）

# problem14
全探索。一度出てきた数値はmapに格納して計算量を削減。
逆探索でできそうでできない問題でつらい。

# problem15
python2.7で実装。
それぞれの格子点に何通りの行き方があるかをラベルし、それを足していく。
ある点の行き方のルート数は、そのルートの上と左のルート数の和になることから。
計算回数はセル数の２乗に比例する。

