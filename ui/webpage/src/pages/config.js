import { Divider } from 'antd';
import styles from './config.css';

export default function() {
  return (
    <div className={styles.normal}>
      <div className={styles.body}>
        <div className={styles.msg}>更多设置选项，请查看使用手册，或私聊机器人发送“设置”查询和使用。</div>
        <div className={styles.inline}>
          <div className={styles.title}>插件主人（超级管理员）QQ：</div>
          <input></input>
          <button>保存</button>
        </div>
        <Divider className={styles.divide}/>
        <div>AA</div>

        <Divider className={styles.divide}/>
        <div>AA</div>

        <Divider className={styles.divide}/>
        <div>AA</div>

        <div>AA</div>
        <Divider className={styles.divide}/>
        <div>AA</div>

        <Divider className={styles.divide}/>
        <div>AA</div>

        <Divider className={styles.divide}/>
        <div>AA</div>

        <Divider className={styles.divide}/>

      </div>
    </div>
  );
}
