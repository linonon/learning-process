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


>%%
>```annotation-json
>{"created":"2023-09-18T07:53:51.563Z","text":"函数的返回不会释放返回值的 memory","updated":"2023-09-18T07:53:51.563Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":59442,"end":59512},{"type":"TextQuoteSelector","exact":"Since allocations from the heap do not go away when a function returns","prefix":"tion in the C standard library. ","suffix":", you can allocate memory for a "}]}]}
>```
>%%
>*%%PREFIX%%tion in the C standard library.%%HIGHLIGHT%% ==Since allocations from the heap do not go away when a function returns== %%POSTFIX%%, you can allocate memory for a*
>%%LINK%%[[#^3dhr46yqdte|show annotation]]
>%%COMMENT%%
>函数的返回不会释放返回值的 memory
>%%TAGS%%
>
^3dhr46yqdte


>%%
>```annotation-json
>{"created":"2023-09-19T09:54:18.884Z","text":"通常给一些全局的變量來使用","updated":"2023-09-19T09:54:18.884Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":60573,"end":60582},{"type":"TextQuoteSelector","exact":"Box::leak","prefix":"eap and explicitly leak it with ","suffix":" to get a 'static reference to i"}]}]}
>```
>%%
>*%%PREFIX%%eap and explicitly leak it with%%HIGHLIGHT%% ==Box::leak== %%POSTFIX%%to get a 'static reference to i*
>%%LINK%%[[#^5w5x9d5qbd6|show annotation]]
>%%COMMENT%%
>通常给一些全局的變量來使用
>%%TAGS%%
>
^5w5x9d5qbd6


>%%
>```annotation-json
>{"created":"2023-09-19T10:02:41.485Z","text":"最常见的 'static , 从一个线程创建的另一个线程不会被这个线程的生命周期约束, 同时另一个线程也不能 refer 任何这个线程的东西, 另一个线程里的任何变量的生命周期只能跟随该线程的生命周期","updated":"2023-09-19T10:02:41.485Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":62774,"end":62792},{"type":"TextQuoteSelector","exact":"std::thread::spawn","prefix":"le of 'static as a bound is the ","suffix":" function that creates a new thr"}]}]}
>```
>%%
>*%%PREFIX%%le of 'static as a bound is the%%HIGHLIGHT%% ==std::thread::spawn== %%POSTFIX%%function that creates a new thr*
>%%LINK%%[[#^lznq38gylb|show annotation]]
>%%COMMENT%%
>最常见的 'static , 从一个线程创建的另一个线程不会被这个线程的生命周期约束, 同时另一个线程也不能 refer 任何这个线程的东西, 另一个线程里的任何变量的生命周期只能跟随该线程的生命周期
>%%TAGS%%
>
^lznq38gylb


>%%
>```annotation-json
>{"created":"2023-09-21T05:18:01.956Z","text":"copy trait 会是按位复制","updated":"2023-09-21T05:18:01.956Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":64823,"end":64841},{"type":"TextQuoteSelector","exact":"copying their bits","prefix":"ate the type’s values simply by ","suffix":". This eliminates all types that"}]}]}
>```
>%%
>*%%PREFIX%%ate the type’s values simply by%%HIGHLIGHT%% ==copying their bits== %%POSTFIX%%. This eliminates all types that*
>%%LINK%%[[#^o35dm1hmx4|show annotation]]
>%%COMMENT%%
>copy trait 会是按位复制
>%%TAGS%%
>
^o35dm1hmx4


>%%
>```annotation-json
>{"created":"2023-09-28T06:20:52.002Z","updated":"2023-09-28T06:20:52.002Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":80673,"end":80730},{"type":"TextQuoteSelector","exact":"It just so happens that that new variable is also named z","prefix":"s only from that point forward. ","suffix":". With that model in mind, this "}]}]}
>```
>%%
>*%%PREFIX%%s only from that point forward.%%HIGHLIGHT%% ==It just so happens that that new variable is also named z== %%POSTFIX%%. With that model in mind, this*
>%%LINK%%[[#^fdny1zg7km4|show annotation]]
>%%COMMENT%%
>
>%%TAGS%%
>
^fdny1zg7km4


>%%
>```annotation-json
>{"created":"2023-09-28T07:33:53.603Z","text":"方差是一个统计量，用于衡量数据分布的离散程度。简单来说，方差越大，数据点越分散；方差越小，数据点越集中。在编程中，方差通常用于评估算法的性能和稳定性。","updated":"2023-09-28T07:33:53.603Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":84242,"end":84360},{"type":"TextQuoteSelector","exact":"Variance is a concept that programmers are often exposed to but rarely know the name of because it’s mostly invisible.","prefix":"eject the code.Lifetime Variance","suffix":" At a glance, variance describes"}]}]}
>```
>%%
>*%%PREFIX%%eject the code.Lifetime Variance%%HIGHLIGHT%% ==Variance is a concept that programmers are often exposed to but rarely know the name of because it’s mostly invisible.== %%POSTFIX%%At a glance, variance describes*
>%%LINK%%[[#^0etjp6q4oqw|show annotation]]
>%%COMMENT%%
>方差是一个统计量，用于衡量数据分布的离散程度。简单来说，方差越大，数据点越分散；方差越小，数据点越集中。在编程中，方差通常用于评估算法的性能和稳定性。
>%%TAGS%%
>
^0etjp6q4oqw


>%%
>```annotation-json
>{"created":"2023-10-16T08:17:41.079Z","text":"字節對齊","updated":"2023-10-16T08:17:41.079Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":91577,"end":91586},{"type":"TextQuoteSelector","exact":"Alignment","prefix":"nd the performance of your code.","suffix":"Before we talk about how a type’"}]}]}
>```
>%%
>*%%PREFIX%%nd the performance of your code.%%HIGHLIGHT%% ==Alignment== %%POSTFIX%%Before we talk about how a type’*
>%%LINK%%[[#^g3mm9mzk1p8|show annotation]]
>%%COMMENT%%
>字節對齊
>%%TAGS%%
>
^g3mm9mzk1p8


>%%
>```annotation-json
>{"created":"2023-10-16T08:21:56.466Z","text":"Rust 提供了 repr(C) -- representation C. 让编译出来的变成 C 的 layout. 在 unsafe 的类别转换上也常用到","updated":"2023-10-16T08:21:56.466Z","document":{"title":"Rust-for-Rustaceans.pdf","link":[{"href":"urn:x-pdf:dcee1bbabcb61f52313b1e32dced5a17"},{"href":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf"}],"documentFingerprint":"dcee1bbabcb61f52313b1e32dced5a17"},"uri":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","target":[{"source":"vault:/learning-process/reading/Rust-for-Rustaceans.pdf","selector":[{"type":"TextPositionSelector","start":95248,"end":95449},{"type":"TextQuoteSelector","exact":"Rust provides a repr attribute you can add to your type definitions to request a particular in-memory rep-resentation for that type. The most common one you will see, if you see one at all, is repr(C).","prefix":"underlying principles. Luckily, ","suffix":" As the name suggests, it lays o"}]}]}
>```
>%%
>*%%PREFIX%%underlying principles. Luckily,%%HIGHLIGHT%% ==Rust provides a repr attribute you can add to your type definitions to request a particular in-memory rep-resentation for that type. The most common one you will see, if you see one at all, is repr(C).== %%POSTFIX%%As the name suggests, it lays o*
>%%LINK%%[[#^rkkl7l7nz9r|show annotation]]
>%%COMMENT%%
>Rust 提供了 repr(C) -- representation C. 让编译出来的变成 C 的 layout. 在 unsafe 的类别转换上也常用到
>%%TAGS%%
>
^rkkl7l7nz9r
