import {Row, Col} from 'antd';


function HomepageContent(props) {
    return (
        <div>
            <Row gutter={[16, 24]}>
                <Col className="gutter-row" span={18}>
                    <Row gutter={[16, 24]}>
                        <Col className="gutter-row" span={24}>
                            <div style={{background: '#0092ff', padding: '8px 0'}}>趋势统计</div>
                        </Col>
                        <Col className="gutter-row" span={12}>
                            <div style={{background: '#0092ff', padding: '8px 0'}}>击杀价值排序</div>
                        </Col>
                        <Col className="gutter-row" span={12}>
                            <div style={{background: '#0092ff', padding: '8px 0'}}>击杀数量排序</div>
                        </Col>
                    </Row>
                </Col>
                <Col className="gutter-row" span={6}>
                    <Row gutter={[16, 24]}>
                        <Col className="gutter-row" span={24}>
                            <div style={{background: '#0092ff', padding: '8px 0'}}>危险星系</div>
                        </Col>
                        <Col className="gutter-row" span={24}>
                            <div style={{background: '#0092ff', padding: '8px 0'}}>危险角色</div>
                        </Col>
                    </Row>
                </Col>
            </Row>
        </div>
    );
}

export default HomepageContent
