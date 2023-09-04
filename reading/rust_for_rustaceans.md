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
