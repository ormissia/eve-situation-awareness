import React, {Component} from "react";
import {Dropdown, Input, Menu} from "semantic-ui-react";
import TabCalendar from "./TabCalendar";

import SSOImg from "../assets/eve-sso-login-black-large.png";

const todayZero = new Date(new Date().setHours(0, 0, 0, 0))
const tomorrowZero = new Date(new Date().setHours(0, 0, 0, 0) + 24 * 60 * 60 * 1000 - 1)
const defaultDateRange = new Map([
    ['hour', [todayZero, tomorrowZero]],
    ['day', [new Date(new Date().setDate(todayZero.getDate() - 30)), tomorrowZero]],
    ['month', [new Date(new Date().setMonth(todayZero.getMonth() - 12)), tomorrowZero]],
    ['year', [new Date(new Date().setFullYear(todayZero.getFullYear() - 3)), tomorrowZero]],
])

export default class TabMenu extends Component {
    state = {
        activeItem: 'day',
        dateRange: defaultDateRange.get('day'),
    }

    handleItemClick = (e, {name}) => {
        this.setState({activeItem: name})
        this.setState({dateRange: defaultDateRange.get(name)})
        const startTS = defaultDateRange.get(name)[0]
        const endTS = defaultDateRange.get(name)[1]
        this.props.getTabMenuTimeType(name, startTS.valueOf(), endTS.valueOf())
    }

    getDateRange = (value) => {
        this.props.getTabMenuTimeType(this.state.activeItem, value[0].valueOf(), value[1].valueOf())
        this.setState({dateRange: value})
    }

    componentDidMount() {
        const startTS = defaultDateRange.get(this.state.activeItem)[0]
        const endTS = defaultDateRange.get(this.state.activeItem)[1]
        this.props.getTabMenuTimeType(this.state.activeItem, startTS.valueOf(), endTS.valueOf())
    }

    render() {
        const {activeItem} = this.state
        return (
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
                        <TabCalendar
                            defaultDateRange={this.state.dateRange}
                            getDateRange={this.getDateRange}
                        />
                    </Menu.Item>
                    <Menu.Item>
                        <Input icon='search' placeholder='Solar System'/>
                    </Menu.Item>
                    <Menu.Item>
                        <Dropdown text='Login'>
                            <Dropdown.Menu>
                                <Dropdown.Item><img src={SSOImg} alt='login'/></Dropdown.Item>
                            </Dropdown.Menu>
                        </Dropdown>
                    </Menu.Item>
                </Menu.Menu>
            </Menu>
        )
    }
}
