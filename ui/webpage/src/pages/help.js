
import styles from './help.css';

export default function() {
  return (
    <div className={styles.normal}>
      <div className={styles.body}>
        <iframe src="/playGuide.html" frameBorder="0" style={{width:"100%", height:"100%"}}></iframe>
      </div>
    </div>
  );
}
