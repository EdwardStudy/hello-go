# Select

select语句属于条件分支流程控制方法，不过它只能用于通道。

select语句中的case关键字只能后跟用于通道的发送操作的表达式以及接收操作的表达式或语句。

在接收操作时，如果在当时有数据的通道多于一个，那么Go语言会通过一种伪随机的算法来决定哪一个case将被执行。

在发送操作时，当有多个case中的通道未满时，它们会被随机选择。