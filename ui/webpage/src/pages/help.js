
import styles from './help.css';
import playGuide from '../assets/playGuide.html';

export default function() {
  return (
    <div className={styles.normal}>
      <div className={styles.body}>
        <iframe title="help" src={playGuide} frameBorder="0" style={{width:"100%", height:"80vh"}}></iframe>
      </div>
    </div>
  );
}
