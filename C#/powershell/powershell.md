powershell相关
=============

报Cloud not load file or assembly System Management.Automation....错误
1. 项目目录下找到 ***.csproj文件
2. 查找Reference Include类似如下代码
```
    <Reference Include="System.Management.Automation, Version=3.0.0.0, Culture=neutral, PublicKeyToken=31bf3856ad364e35, processorArchitecture=MSIL">
      <HintPath>..\packages\Microsoft.PowerShell.5.ReferenceAssemblies.1.1.0\lib\net4\System.Management.Automation.dll</HintPath>
    </Reference>
```
3. 修改为：
```
 <Reference Include="System.Management.Automation" />
```