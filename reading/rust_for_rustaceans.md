annotation-target:: Rust-for-Rustaceans.pdf


>%%
>```annotation-json
>{"created":"2023-08-21T07:06:45.748Z","text":"記憶體術語","updated":"2023-08-21T07:06:45.748Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":48163,"end":48181},{"type":"TextQuoteSelector","exact":"Memory Terminology","prefix":"at they are worth covering here.","suffix":"Before we dive into regions of m"}]}]}
>```
>%%
>*%%PREFIX%%at they are worth covering here.%%HIGHLIGHT%% ==Memory Terminology== %%POSTFIX%%Before we dive into regions of m*
>%%LINK%%[[#^wxhijpsierp|show annotation]]
>%%COMMENT%%
>記憶體術語
>%%TAGS%%
>
^wxhijpsierp



>%%
>```annotation-json
>{"created":"2023-09-04T15:54:14.724Z","text":"生命周期相关, 值被移动了以后, 就和以前的变量没关系了","updated":"2023-09-04T15:54:14.724Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":52138,"end":52412},{"type":"TextQuoteSelector","exact":" When a variable is later accessed, you can imagine drawing a line from the previous access of that variable to the new access, which establishes a dependency relation-ship between the two accesses. If the value in a variable is moved, no lines can be drawn from it anymore.","prefix":" then on named by that variable.","suffix":"In this model, a variable exists"}]}]}
>```
>%%
>*%%PREFIX%%then on named by that variable.%%HIGHLIGHT%% ==When a variable is later accessed, you can imagine drawing a line from the previous access of that variable to the new access, which establishes a dependency relation-ship between the two accesses. If the value in a variable is moved, no lines can be drawn from it anymore.== %%POSTFIX%%In this model, a variable exists*
>%%LINK%%[[#^fi8rfs3gic6|show annotation]]
>%%COMMENT%%
>生命周期相关, 值被移动了以后, 就和以前的变量没关系了
>%%TAGS%%
>
^fi8rfs3gic6


>%%
>```annotation-json
>{"created":"2023-09-05T14:41:51.350Z","text":"阴影做法, 上一个相同变量名的变量, 是这个变量的影子, 实体是这个变量","updated":"2023-09-05T14:41:51.350Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":54622,"end":54691},{"type":"TextQuoteSelector","exact":"shadowing —the later variable “shadows” the former by the same name. ","prefix":"tinct variables. This is called ","suffix":"The two variables coexist, thoug"}]}]}
>```
>%%
>*%%PREFIX%%tinct variables. This is called%%HIGHLIGHT%% ==shadowing —the later variable “shadows” the former by the same name.== %%POSTFIX%%The two variables coexist, thoug*
>%%LINK%%[[#^jqaejw01m7|show annotation]]
>%%COMMENT%%
>阴影做法, 上一个相同变量名的变量, 是这个变量的影子, 实体是这个变量
>%%TAGS%%
>
^jqaejw01m7


>%%
>```annotation-json
>{"created":"2023-09-12T14:18:05.048Z","text":"最重要的三个内存区域是堆栈、堆和静态内存","updated":"2023-09-12T14:18:05.048Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":57230,"end":57269},{"type":"TextQuoteSelector","exact":"the stack, the heap, and static memory.","prefix":"poses of writ-ing Rust code are ","suffix":"The StackThe stack is a segment "}]}]}
>```
>%%
>*%%PREFIX%%poses of writ-ing Rust code are%%HIGHLIGHT%% ==the stack, the heap, and static memory.== %%POSTFIX%%The StackThe stack is a segment*
>%%LINK%%[[#^j75xr8aje1p|show annotation]]
>%%COMMENT%%
>最重要的三个内存区域是堆栈、堆和静态内存
>%%TAGS%%
>
^j75xr8aje1p


>%%
>```annotation-json
>{"created":"2023-09-12T14:20:46.506Z","text":"堆栈顶部分配一个连续的内存块，称为帧","updated":"2023-09-12T14:20:46.506Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":57441,"end":57482},{"type":"TextQuoteSelector","exact":"frame is allocated at the top of the stac","prefix":"iguous chunk of memory called a ","suffix":"k.  Near the bottom of the stack"}]}]}
>```
>%%
>*%%PREFIX%%iguous chunk of memory called a%%HIGHLIGHT%% ==frame is allocated at the top of the stac== %%POSTFIX%%k.  Near the bottom of the stack*
>%%LINK%%[[#^ibhrzi1trjb|show annotation]]
>%%COMMENT%%
>堆栈顶部分配一个连续的内存块，称为帧
>%%TAGS%%
>
^ibhrzi1trjb


>%%
>```annotation-json
>{"created":"2023-09-12T14:21:04.823Z","text":"在 stack 的底部","updated":"2023-09-12T14:21:04.823Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":57536,"end":57549},{"type":"TextQuoteSelector","exact":"main function","prefix":" the stack is the frame for the ","suffix":", and as functions call other fu"}]}]}
>```
>%%
>*%%PREFIX%%the stack is the frame for the%%HIGHLIGHT%% ==main function== %%POSTFIX%%, and as functions call other fu*
>%%LINK%%[[#^u17d8rngb6|show annotation]]
>%%COMMENT%%
>在 stack 的底部
>%%TAGS%%
>
^u17d8rngb6


>%%
>```annotation-json
>{"created":"2023-09-12T14:22:11.519Z","text":"存储在堆栈帧上的任何变量在帧消失后都不能被访问, 这点和 Rust 的 LifeTime 概念十分相似","updated":"2023-09-12T14:22:11.519Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":58263,"end":58315},{"type":"TextQuoteSelector","exact":"very closely tied to the notion of lifetimes in Rust","prefix":" they eventually disappear, are ","suffix":". Any variable stored in a frame"}]}]}
>```
>%%
>*%%PREFIX%%they eventually disappear, are%%HIGHLIGHT%% ==very closely tied to the notion of lifetimes in Rust== %%POSTFIX%%. Any variable stored in a frame*
>%%LINK%%[[#^sghjpy5xfa|show annotation]]
>%%COMMENT%%
>存储在堆栈帧上的任何变量在帧消失后都不能被访问, 这点和 Rust 的 LifeTime 概念十分相似
>%%TAGS%%
>
^sghjpy5xfa


>%%
>```annotation-json
>{"created":"2023-09-18T07:48:31.680Z","text":"给函数的返回值预先开辟空间","updated":"2023-09-18T07:48:31.680Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":58778,"end":58947},{"type":"TextQuoteSelector","exact":"If that value is the function’s return value, the calling function can leave some space on its stack for the called function to write that value into before it returns. ","prefix":" the cur-rent function’s frame. ","suffix":"But if you want to, say, send th"}]}]}
>```
>%%
>*%%PREFIX%%the cur-rent function’s frame.%%HIGHLIGHT%% ==If that value is the function’s return value, the calling function can leave some space on its stack for the called function to write that value into before it returns.== %%POSTFIX%%But if you want to, say, send th*
>%%LINK%%[[#^m5jasdksdvb|show annotation]]
>%%COMMENT%%
>给函数的返回值预先开辟空间
>%%TAGS%%
>
^m5jasdksdvb
