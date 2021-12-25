import {BrowserRouter as Router, Route, Link, Routes} from 'react-router-dom';
import Homepage from "./pages/Homepage";

function App() {
    return (
        <Router>
            <div className='App'>
                <Routes>
                    <Route path='/' element={<Homepage/>}/>
                </Routes>
            </div>
        </Router>
    );
}

export default App;
