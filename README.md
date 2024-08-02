## コマンド

- `lc init`           : lc を使用するための準備をする｡ディレクトリの生成やファイル生成､値の設定を行う｡
  - `--version` : 使用している vscode-leetcode のバージョンを設定する｡

- `lc new-problem`    : 新たに問題を解くための準備をする｡必要なディレクトリやファイルの生成､値の設定を行う｡

- `lc change-problem` : 解く問題を変更する｡

- `lc update-step`    : stepファイルをナンバリングする｡

- `lc generate-md`    : stepファイルを集約したMarkdownファイルを生成する｡

## 使い方

### 初期設定

**1. 任意のディレクトリで `lc init --path "<vscode-leeetcodeのバージョン>"` を実行する｡**
  
  1.  カレントディレクトリに `leetcode/`, `config.json`, `problems/` が生成される
  
  2.  `config.json` > `"leetcodeDirPath"` に ルートから`leetcode/` までのパスが設定される｡
  
  3.  `config.json` > `"vscodeLeetcodeVersion"` に `--path` の値が設定される｡
  
  4.  vscode-leetcode ディレクトリにファイルが生成される｡

### 問題を解く

**2. `<ワークブック名>-<問題番号>-<問題名>` という名前で origin main からブランチを切って､ブランチを切り変える｡**

**3. `leetcode/` の配下で `lc new-problem` を実行する｡**
  
  1. `problems/` 配下に `<ワークブック名>/<問題番号><問題名>`が生成される｡
  
  2. `<問題番号><問題名>/` 配下に `step_count.json` が生成される｡
  
  3. `leetcode/config.json` > `"problemDirPath"` に `<問題番号><問題名>/` までのパスが設定される｡

**4. VSCode で `Code Now` を押下し､`step_x.py`を生成し､問題を解く｡**

### 解いた問題を管理する

**5. `leetcode/<ワークブック名>/<問題番号><問題名>/` の配下で `lc update-step` を実行する｡**
  
  1. `step_x.py` が `step_count.json`> `"stepNumber"` をもとにナンバリングされる｡
  
  2. `step_count.json`> `"stepNumber"` が加算される｡
  
### Markdownを生成する

**6. `leetcode/<ワークブック名>/<問題番号><問題名>/` の配下で `lc generate-md` を実行する｡**
  
  1. `step_<number>.py` が集約された `problem.md` が生成される｡

-----------------------

## 用語

- ワークブック    : LeetCodeの問題がまとめられた問題リストのこと｡(例 : Arai60 [https://leetcode.com/list/xo2bgr0r/])

- 問題ディレクトリ : ワークブックに含まれる問題ごとのディレクトリ｡

- stepファイル   : 解法が書かれているソースコード｡

## 想定しているディレクトリ構造

```txt
.
└── leetcode
    ├── config.json
    └── problems
        └── <workbook>
            └── <problem>
                ├── problem.md
                ├── step_x.py
                ├── step_1.py
                └── step_count.json
```
          
- `leetcode/`   : 基準となるディレクトリ｡この中ですべてを管理する｡

- `config.json` : lc がうまく動作するための情報を管理する｡
  
  - `"leetcodeDirPath"`       : leetcode/ までのパス｡
  
  - `"problemDirPath"`        : vscode-leetcode が step_x.py を作成するパス｡
  
  - `"vscodeLeetcodeVersion"` : 使用している vscode-leetcode のバージョン｡vscode-leetcode ディレクトリを特定するために使用される｡

- `problems/`   : すべてのワークブックを格納する｡

- `<workbook>`  : ワークブック｡ (例 : Arai 60) 

- `<problem>`   : 問題ディレクトリ｡stepファイルを格納する｡
  
  - `step_x.py`        : ナンバリング前のstepファイル｡
  
  - `step_<number>.py` : ナンバリング後のstepファイル｡(例 : step_1.py)
  
  - `problem.md`       : stepファイルの内容を集約したMarkdownファイル｡
  
  - `step_count.json`  : ナンバリングのための番号が管理されている｡