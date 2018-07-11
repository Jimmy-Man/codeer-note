<?php
/**
 * --------------------------------------
 * Author Jimmy
 * Email  1877876812@qq.com
 * --------------------------------------
 * Zencart change configration
 * Zencart 批量修改配置文件脚本
 * 
 * --------------------------------------
 */

echo "-------Start-----------";
//wwwroot Path 
$web_path = '/home/wwwroot';
//admin path
$admin_path = 'adminfff';
//setting file
$set_file = 'setting.txt';

/**
 * ----------------------------------------------------
 * Format:
 * ----------------------------------------------------
 * domain,databaseName
 * 
 * baidu.com,baidu
 * 
 * 
 * 
 * ----------------------------------------------------
 */
$file = file($set_file);

if(!empty($file)){
    foreach($file as $set){                      
        $s = explode(',',$set);
        $domain = trim($s[0]);
        $database = trim($s[1]);

        $parge_arr = array(
            '/\'BASE_PATH\'\W*,\W*\'(.*)\'/',
            '/\'DB_DATABASE\',\W*\'(.*)\'/'
        );
        $replae_arr = array(
            "'BASE_PATH','{$web_path}/{$domain}/'",
            "'DB_DATABASE', '{$database}'"
        );

        $include_conf = $web_path.'/'.$domain.'/includes/configure.php';
        if(file_exists($include_conf)){
            $setting_text = file_get_contents($include_conf);        
            $new_setting  = preg_replace($parge_arr,$replae_arr,$setting_text);
            file_put_contents($include_conf,$new_setting);
        }else{
            echo $include_conf.'File not Exists';
        }

        $admin_conf = $web_path.'/'.$domain.'/'.$admin_path.'/includes/configure.php';
        if(file_exists($admin_conf)){
            $setting_text = file_get_contents($admin_conf);        
            $new_setting  = preg_replace($parge_arr,$replae_arr,$setting_text);
            file_put_contents($admin_conf,$new_setting);
        }else{
            echo $admin_conf.'File not Exists';
        }
    }
}
echo "-------End---------------";

?>
