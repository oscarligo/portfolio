
// ThemeStort is a simple store to manage the theme of the application. 
// It uses a private state variable to track whether the dark mode is 
// enabled or not, and provides a method to toggle the theme. 
class ThemeStore {
    
    // Private state variable to track the theme
    #isDark = $state(false);

    // Getter to access the current theme state
    get isDark() {
        return this.#isDark;
    }

    // Method to toggle the theme
    toggleTheme() {
        this.#isDark = !this.#isDark;
        if (this.#isDark) {
            document.documentElement.classList.add('dark');
        } else {
            document.documentElement.classList.remove('dark');
        }
    }
}

export const theme = new ThemeStore();