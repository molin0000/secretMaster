
import styles from './about.css';
import { Component } from 'react';
import { Card } from 'antd';
import { apiGet } from './utils/utils.js';

class About extends Component {
  constructor(props) {
    super(props);
    this.state = { version: 'v3.2.4' }
  }

  async componentWillMount() {
    let ret = await apiGet('version');
    console.log(ret);
    this.setState({ version: ret.data.data.Version });
  }

  render() {
    return (
      <div className={styles.body}>
        <Card style={{maxWidth:"700px"}}>
          <h2 style={{ marginLeft: '15px' }}>当前版本：{this.state.version}</h2>
          <div className={styles.text}>欢迎使用序列战争插件~O(∩_∩)O~</div>
          <div className={styles.text}>这是一款《诡秘之主》背景的QQ群游戏</div>
          <div className={styles.text}>在群友们的共同努力下，现已经开发出了许多独具魅力的功能系统</div>
          <div className={styles.text}>有问题欢迎到酷Q论坛提问：<a href='https://cqp.cc/t/46674' rel="noopener noreferrer" target="_blank">https://cqp.cc/t/46674</a></div>
          <div className={styles.text}>有好的想法也欢迎来策划QQ群：1028799086</div>
          <div className={styles.text}>想直接游玩可以来游戏QQ群：1030551041 </div>
          <div className={styles.text}>以及《诡秘之主》粉丝序列群:731419992</div>
          <div className={styles.text}>支持自定义文字副本，副本编辑器地址：<a href='https://mission-editor.now.sh/' rel="noopener noreferrer" target="_blank">https://mission-editor.now.sh</a></div>
          <div className={styles.text}>本插件基于Coolq Go SDK开发，代码完全开源，欢迎共同学习和交流</div>
          <div className={styles.text}>开源地址：<a href='https://github.com/molin0000/secretMaster' rel="noopener noreferrer" target="_blank">https://github.com/molin0000/secretMaster</a></div>
          <div className={styles.text}>如果喜欢，请给我发电：<a href='https://afdian.net/@molin' rel="noopener noreferrer" target="_blank">https://afdian.net/@molin</a></div>
          <div className={styles.text}>
            <span role="img" aria-label="heart">❤️</span>O(∩_∩)O<span role="img" aria-label="heart">❤️</span>空想之喵<span role="img" aria-label="heart">❤️</span>
          </div>
        </Card>
      </div>
    );
  }
}

export default About;