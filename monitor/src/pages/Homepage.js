import {
    Header,
    Container,
    Input,
    Menu,
    Segment,
    Dropdown,
} from 'semantic-ui-react'
import SSOImg from '../assets/eve-sso-login-black-large.png'

import React, {Component} from "react";
import axios from "axios";
import Chart from "../components/Chart";

const style = {
    h1: {
        padding: '2em',
    },
    h2: {
        margin: '4em 0em 2em',
    },
    h3: {
        marginTop: '2em',
        padding: '2em 0em',
    },
    last: {
        marginBottom: '300px',
    },
}

class Homepage extends Component {

    state = {
        activeItem: 'day',
        chartOption: {
            title: {
                text: '击杀趋势',
                left: 'center',
                bottom: '0',
            },
            tooltip: {},
            xAxis: {
                data: []
            },
            yAxis: [
                {
                    splitLine: {show: false},
                    type: 'value',
                    name: 'kill quantity',
                    splitNumber: 5, //设置坐标轴的分割段数
                },
                {
                    splitLine: {show: false},
                    type: 'value',
                    name: 'kill value',
                    splitNumber: 5, //设置坐标轴的分割段数
                }
            ],
            series: [
                {
                    name: 'kill quantity',
                    type: 'line',
                    smooth: true,
                    yAxisIndex: 0,
                    data: []
                },
                {
                    name: 'kill value',
                    type: 'bar',
                    yAxisIndex: 1,
                    data: []
                }
            ]
        }
    }

    handleItemClick = (e, {name}) => {
        this.setState({activeItem: name})
        console.log(name)
        axios.get('http://localhost:8080/web/solar_system_kill?time_type='+name+'&start_time_stamp=1639929600000&end_time_stamp=&page_size=24')
            .then(response => {
                console.log(response.data.data.Y.kill_quantity)
                const initOption = {
                    xAxis: {
                        data: response.data.data.X
                    },
                    series: [
                        {
                            name: 'kill quantity',
                            type: 'line',
                            smooth: true,
                            yAxisIndex: 0,
                            data: response.data.data.Y.kill_quantity
                        },
                        {
                            name: 'kill value',
                            type: 'bar',
                            yAxisIndex: 1,
                            data: response.data.data.Y.kill_value
                        }
                    ]
                }
                const newOption = Object.assign(this.state.chartOption, initOption)
                console.log(111)
                console.log(newOption)
                this.setState({chartOption: newOption})
                console.log(222)
                console.log(this.state.chartOption)
            })
    }

    render() {
        const {activeItem} = this.state

        return (
            <div>
                <Container>
                    <Menu pointing>
                        <Menu.Item
                            name='year'
                            active={activeItem === 'year'}
                            onClick={this.handleItemClick}
                        />
                        <Menu.Item
                            name='month'
                            active={activeItem === 'month'}
                            onClick={this.handleItemClick}
                        />
                        <Menu.Item
                            name='day'
                            active={activeItem === 'day'}
                            onClick={this.handleItemClick}
                        />
                        <Menu.Item
                            name='hour'
                            active={activeItem === 'hour'}
                            onClick={this.handleItemClick}
                        />
                        <Menu.Menu position='right'>
                            <Menu.Item>
                                <Input icon='search' placeholder='Solar System'/>
                            </Menu.Item>
                            <Menu.Item>
                                <Dropdown text='Login'>
                                    <Dropdown.Menu>
                                        <Dropdown.Item><img src={SSOImg}/></Dropdown.Item>
                                    </Dropdown.Menu>
                                </Dropdown>
                            </Menu.Item>
                        </Menu.Menu>

                    </Menu>

                    <Header as='h1' content='EVE Situation Awareness' style={style.h1} textAlign='center'/>

                    <Segment>
                        <Chart chartOption={this.state.chartOption}/>
                    </Segment>
                </Container>
            </div>
        )
    }
}


export default Homepage
