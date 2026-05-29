# Atlas

## aseprite
- aseprite 的frame包含 duration 字段，用于描述帧的持续时间。默认值为 0.1 秒，而atlas 没有
- aseprite 的meta包含FrameTags 字段，用于描述帧的标签。
- aseprite 的frame.frame 没有idx字段，而atlas 有，atlas 中的idx字段是帧的索引，从0开始。

## 参考
- layaair atlas:https://github.com/layabox/LayaAir/blob/LayaAir_3.3/src/layaAir/laya/loaders/AtlasLoader.ts
- spine atlas
- cocos plist
- aseprite
