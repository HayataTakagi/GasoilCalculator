# GasoilCalculator
- VOYAGE GROUPのインターン[Treasure](https://voyagegroup.com/internship/treasure/)の中間課題として作成

## 交通費(車)計算サービス
### 概要
- 距離を入力するだけで交通費が分かる
- ガソリン代(¥/l)はAPIから取得
- 燃費は代表的な車から選択式にする
- CO2排出量を表示する
- CO2排出量を何かに置換して表示する(木のが一年間に吸収する量等)

## バックエンド
- [x] ガソリン価格をスクレイピングし、フロントに返却
- [x] 燃費データをDBに保存しておき、フロントに燃費データを返却

![Screen Shot 2019-08-19 at 2 49 40](https://user-images.githubusercontent.com/28585609/63228332-282e8e00-c22c-11e9-8487-128428efe912.png)
![Screen Shot 2019-08-19 at 2 49 57](https://user-images.githubusercontent.com/28585609/63228336-2d8bd880-c22c-11e9-94bf-8a0537972bc2.png)
![Screen Shot 2019-08-19 at 2 50 08](https://user-images.githubusercontent.com/28585609/63228337-311f5f80-c22c-11e9-8fbc-dcf3f5d5a522.png)


### local環境
```
❯ docker --version
Docker version 18.09.2, build 6247962

❯ go version
go version go1.12.7 darwin/amd64

❯ node --version
v10.15.3

❯ npm -v
6.4.1
```

### Links
 - https://12factor.net
