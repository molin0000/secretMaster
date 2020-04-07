import styles from './index.css';
import { Component } from 'react';
import { Button, Statistic, Row, Col, Input, Divider } from 'antd';
const { TextArea } = Input;

class Home extends Component {
  render() {
    return (
      <div className={styles.body}>
        <Row gutter={16}>
          <Row>
            <Col span={6}>
              <Button type='primary' style={{width:'20vw', marginTop:'10px'}}>QQ登录</Button>
            </Col>
            <Col span={9}>
              <Statistic title="当前QQ号码" value={1234567890} groupSeparator={""} />
            </Col>
            <Col span={9}>
              <Statistic title="昵称" value={"空想之喵"} />
            </Col>
          </Row>
          <Row>
            <Col span={1}></Col>
            <Col span={22}>
              <TextArea rows={16} style={{ marginTop:"10px", marginBottom:"10px", background: "#660066", borderRadius: "20px", color: "white"}} value={`To: 空想之喵

昵称：空想之喵
途径：魔女
序列：序列4：绝望
勋章：3🎖🎖🎖
经验：6343
金镑：13
幸运：0
灵力：200
修炼时间：857小时
战力评价：超凡入圣
教会/组织：喵喵会
工作：无业游民
尊名：无`}>

              </TextArea>
            </Col>
            <Col span={1}></Col>
          </Row>
          <Row>
            <Row>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>探险</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>钓鱼</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>副本</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>许愿</Button></Col>
            </Row>
            <Row>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>速算</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>学识</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>阵营</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>排行</Button></Col>
            </Row>
            <Row>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>祈祷</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>商店</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>属性</Button></Col>
              <Col span={6}><Button type='primary' style={{marginBottom:"10px", width:"20vw"}}>排行</Button></Col>
            </Row>
          </Row>
          <Divider style={{width:"80%"}}/>
          <Row>
            <Col span={2}></Col>
            <Col span={16}><Input/></Col>
            <Col span={4}><Button type='primary' style={{width:"20vw", marginLeft:"10px"}}>发送</Button></Col>
          </Row>
        </Row>
      </div>
    );
  }
}

export default Home;
