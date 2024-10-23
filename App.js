Import react, { component } from 'react'
import { 
    View, StatusBar, Text
} from 'react-native';

class App extends Component {
    constructor(props) {
        super(props);
        this.state = { 
         hasil: 0,
         hitung: 0,
        };
    }
    render() {
        return (
        <View style={{flex: 1, backgroundColor: '#212121'}}>
         <StatusBar backgroundColor="#212121" barStyle="light-content" />        
         
         <View style={{flexDirection: 'row', }}>
            <view style={{flex: 1}}>
                <text style={{color: '#FFFFFF',  fontSize: 24}}>7</text>
            </view>
            <view style={{flex: 1}}>
                <text style={{color: '#FFFFFF',  fontSize: 24}}>8</text>
            </view>
            <view style={{flex: 1}}>
                <text style={{color: '#FFFFFF',  fontSize: 24}}>9</text>
            </view>

         </View>
         
         </View>

        );
    }
}

export default App;