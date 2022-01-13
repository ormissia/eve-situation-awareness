import React, {Component} from "react";
import {Layout} from 'antd';
import moment from "moment";

import HomepageContent from "../components/HomepageContent";
import HomepageHeader from "../components/HomepageHeader";

const {Header, Footer, Content} = Layout;

class Homepage extends Component {

    state = {
        conditions: {
            group: 'day',
            startTime: moment().startOf('day').add(1, 'days').subtract(1, 'weeks').valueOf(),
            endTime: moment().valueOf(),
            solarSystem: '',
        }
    }

    headerCallback(conditions) {
        // console.log(conditions)
        this.setState(conditions)

        // 通过子组件的实例调用组组件中的方法
        this.childRef.loadData()
    }

    handleChildEvent = (ref) => {
        // 将子组件的实例存到 this.childRef 中, 这样整个父组件就能拿到
        this.childRef = ref
    }

    render() {
        return (
            <Layout style={{background: 'white'}}>
                <Header style={{position: 'fixed', zIndex: 1, width: '100%', background: 'whitesmoke'}}>
                    <HomepageHeader
                        initConditions={this.state.conditions}
                        callback={this.headerCallback.bind(this)}
                    />
                </Header>
                <Content className="site-layout" style={{padding: '32px', marginTop: 64}}>
                    <HomepageContent
                        onChildEvent={this.handleChildEvent}
                        conditions={this.state.conditions}
                    />
                </Content>
                <Footer style={{background: 'whitesmoke'}}>
                    foooooot
                </Footer>
            </Layout>
        )
    }
}


export default Homepage
