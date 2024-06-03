## そもそも生成AIとは
- https://aws.amazon.com/jp/what-is/generative-ai/
### どんなことができるか
#### ※上記リンクの写経
- 会話や画像、動画、音楽の生成が得意
- 画像認識や自然言語処理とは異なる
- 例えば、チャットボットやメディア制作ができる

- 生産性の向上が見込める
　- 研究の加速
    - パターン探したりとかができる
    - ex) 製薬業界で創薬を加速
- カスタマーエクスペリエンス改善
- 従業員の生産性向上
### 仕組み
- 基盤モデル
- 大規模言語モデル
## AWSにはどのような生成AIサービスが存在するか
- Amazon Q Developer
  - AWSで作業するときのアシスタント
    - `Amazon Q Developer は、AWS アプリケーションの理解、構築、拡張、運用を支援する、生成型人工知能 (AI) を搭載した会話型アシスタントです。AWS アーキテクチャ、AWS リソース、ベストプラクティス、ドキュメント、サポートなどについて質問できます。Amazon Q は機能を継続的に更新しているため、質問に対して最も状況に即した実用的な回答が得られます。`
    - https://docs.aws.amazon.com/amazonq/latest/qdeveloper-ug/what-is.html#:~:text=Amazon%20Q%20Developer%20is%20a,extend%2C%20and%20operate%20AWS%20applications.
  - Amazon Bedrock上に構築されている
- CodeWhisperer
  - コーディングのお手伝いAI
    - https://aws.amazon.com/jp/codewhisperer/
- Amazon Bedrock
  - いろいろできるサービス
    - 画像生成, テキスト生成
  - https://aws.amazon.com/jp/bedrock/
- Amazon SageMaker JumpStart
  - 機械学習
    - 事前にトレーニングされたモデルを使用できる
  - https://docs.aws.amazon.com/ja_jp/sagemaker/latest/dg/studio-jumpstart.html
- AWS HealthScribe
  - 医療系
  - 患者と医者の会話をテキストに起こしたり
  - https://docs.aws.amazon.com/ja_jp/transcribe/latest/dg/health-scribe.html

- QuickSight
  - BIツール
  - https://aws.amazon.com/jp/quicksight/?amazon-quicksight-whats-new.sort-by=item.additionalFields.postDateTime&amazon-quicksight-whats-new.sort-order=desc
