import { Divider, Button, Switch, Table, Select } from 'antd';
import styles from './config.css';

const { Option } = Select;

const groupColumns = [
  {
    title: '序号',
    dataIndex: 'key',
    key: 'key',
  },
  {
    title: '群号',
    dataIndex: 'group',
    key: 'group',
  },
  {
    title: '人数',
    dataIndex: 'amount',
    key: 'amount',
  },
  {
    title: '管理员',
    dataIndex: 'master',
    key: 'master',
  },
  {
    title: '操作',
    dataIndex: 'operation',
    key: 'operation',
    render: (text, record) => (
      <span>
        <a>退群</a>
      </span>
    ),
  },
  {
    title: '开关',
    dataIndex: 'switch',
    key: 'switch',
    render: (text, record) => (
      <span>
        <Switch checkedChildren="开" unCheckedChildren="关" defaultChecked />
      </span>
    ),
  },
]

const groupData = [
  {
    key: '1',
    group: '132423442',
    amount: '132/200',
    master: '234234234',
  },
  {
    key: '2',
    group: '132423442',
    amount: '132/200',
    master: '234234234',
  },
  {
    key: '3',
    group: '132423442',
    amount: '132/200',
    master: '234234234',
  }
]

export default function () {
  return (
    <div className={styles.normal}>
      <div className={styles.body}>
        <div className={styles.msg}>更多设置选项，请查看使用手册，或私聊机器人发送“设置”查询和使用。</div>
        <div className={styles.inline}>
          <div className={styles.title}>插件主人（超级管理员）QQ：</div>
          <input className={styles.input} />
          <Button type="primary">保存</Button>
        </div>
        <Divider className={styles.divide} />
        <div className={styles.inline}>
          <div className={styles.title}>消息回复延迟（毫秒）：</div>
          <input className={styles.input} style={{ marginLeft: "35px" }} />
          <Button type="primary">保存</Button>
        </div>
        <Divider className={styles.divide} />
        <div className={styles.inline}>
          <div className={styles.title}>附件扫描：</div>
          <div className={styles.title} style={{ marginLeft: "90px" }}>副本：16 个， 学识：173 个，宠物：120 个，宠物投食：220 个</div>
        </div>
        <Divider className={styles.divide} />
        <div className={styles.inline}>
          <div className={styles.title}>分群管理：</div>
          <div className={styles.text}>全局开关：</div>
          <Switch checkedChildren="开" unCheckedChildren="关" defaultChecked />
          <Button type="primary" style={{ marginLeft: "360px" }}>保存</Button>
        </div>
        <Table columns={groupColumns} dataSource={groupData} size="small" />
        <Divider className={styles.divide} />
        <div className={styles.inline}>
          <div className={styles.title}>货币映射：</div>
          <div className={styles.text}>启用：</div>
          <Switch checkedChildren="开" unCheckedChildren="关" defaultChecked />
          <Button type="primary" style={{ marginLeft: "380px" }}>保存</Button>
        </div>
        <div className={styles.inline}>
          <div className={styles.text2}>ini文件路径：</div>
          <input className={styles.input} style={{ margin: "4px 0px 4px 84px" }} />
        </div>
        <div className={styles.inline}>
          <div className={styles.text2}>ini节点必须为QQ号：[QQ]</div>
        </div>
        <div className={styles.inline}>
          <div className={styles.text2}>关键字：</div>
          <input className={styles.input} style={{ margin: "4px 0px 4px 110px", width: "200px"}} />
          <div className={styles.text2} style={{ marginLeft: "80px"}}>编码：</div>
          <Select defaultValue="GB2312" style={{ width: 124, margin:"4px 0" }}>
            <Option value="GB2312">GB2312</Option>
            <Option value="UTF8">UTF8</Option>
          </Select>
        </div>
        <Divider className={styles.divide} />
        <div>AA</div>

        <Divider className={styles.divide} />
        <div>AA</div>

        <Divider className={styles.divide} />

      </div>
    </div>
  );
}
