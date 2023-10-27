import logo from './logo.svg';
import './App.css';
import TabBox from './components/TabBox.tsx';

function App() {
  return (
    <div className="App">
      <div>
        <TabBox 
        numTabs={3} 
        tabNames={["overview", "budget", "transactions", "categories"]}
        />
      </div>
    </div>
  );
}

export default App;
