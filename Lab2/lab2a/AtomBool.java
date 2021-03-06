package lab2a;

public class AtomBool {
	private boolean value;
	
	public AtomBool(boolean value) {
		this.value = value;
	}
	
	public synchronized boolean get() {
		return value;
	}
	
	public synchronized boolean compareAndSet(boolean expect, boolean update) {
		if(value == expect) {
			value = update;
			return true;
		}
		return false;
	}
}
