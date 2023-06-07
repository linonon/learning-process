# Windows 11 设置开机启动执行脚本

要在Windows 11启动时默认以管理员模式执行"Start-Service sshd"命令，您可以按照以下步骤进行设置：

1. 打开开始菜单，然后搜索"PowerShell"。
2. 右键单击"Windows PowerShell"图标，并选择"以管理员身份运行"。
3. 在提升的权限的PowerShell窗口中，运行以下命令：

   ```powershell
   Set-ExecutionPolicy RemoteSigned -Scope LocalMachine
   ```

   此命令将更改执行策略，以允许在本地计算机范围内运行本地脚本。

4. 创建一个启动脚本文件。在提升的权限的PowerShell窗口中，运行以下命令：

   ```powershell
   New-Item -Path "C:\Windows\System32\GroupPolicy\Machine\Scripts\Startup" -ItemType Directory -Force
   ```

   这将在"启动"文件夹中创建一个新的文件夹以存放启动脚本。

5. 编辑启动脚本文件。在提升的权限的PowerShell窗口中，运行以下命令：

   ```powershell
   New-Item -Path "C:\Windows\System32\GroupPolicy\Machine\Scripts\Startup\start_sshd.ps1" -ItemType File -Force
   notepad "C:\Windows\System32\GroupPolicy\Machine\Scripts\Startup\start_sshd.ps1"
   ```

   这将创建一个名为"start_sshd.ps1"的新文件，并使用记事本打开它以进行编辑。

6. 在记事本中，输入以下命令并保存更改：

   ```powershell
   Start-Process powershell -Verb runAs -ArgumentList 'Start-Service sshd'
   ```

   这将在启动时以管理员权限运行一个新的PowerShell进程，并执行"Start-Service sshd"命令。

7. 关闭记事本。

8. 配置组策略。在提升的权限的PowerShell窗口中，运行以下命令：

   ```powershell
   gpedit.msc
   ```

   这将打开本地组策略编辑器。

9. 在组策略编辑器中，导航到以下路径：

   `计算机配置 -> Windows 设置 -> 脚本 (启动/关机)`

10. 右键单击"启动"，然后选择"属性"。

11. 在"启动"属性对话框中，切换到"脚本"选项卡。

12. 单击"添加"按钮，并选择之前创建的"start_sshd.ps1"脚本文件。

13. 单击"确定"保存更改。

现在，当您启动Windows 11时，"Start-Service sshd"命令将以管理员模式自动执行。

请注意，以上步骤需要管理员权限才能执行。
