import React, {Component} from "react";
import {Dropdown, Input, Menu} from "semantic-ui-react";
import SSOImg from "../assets/eve-sso-login-black-large.png";

export default class TabMenu extends Component {
    state = {
        activeItem: 'day',
    }
    handleItemClick = (e, {name}) => {
        this.setState({activeItem: name})
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
        )
    }
}
