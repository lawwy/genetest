# genetest

《复杂》一书中关于遗传算法与元胞自动机的两个实验，“扫地机器人”实验，与“多数分类”实验。

测试函数中有进化表现不错的基因序列，可通过测试命令观察其任务执行过程

“扫地机器人” 任务测试：
``go test -timeout 60s -run ^Test_Show$ genetest/gene -v``

“多数分类” 任务测试：
``go test -timeout 60s -run ^Test_ShowCell$ genetest/gene -v``
