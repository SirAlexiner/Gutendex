$(document).ready(function () {
    /* Function to handle normal search
    retrieve language and search query
    go to the constructed href */
    function handleSearch() {
        var language = $('#language-select').val();
        var searchQuery = $('#search-input').val();
        window.location.href = '/?languages=' + language + '&search=' + searchQuery;
    }

    // Click event for search button
    $('#search-button').click(function () {
        handleSearch();
    });

    // Enter key event for search input
    $('#search-input').on('keydown', function (event) {
        if (event.key === 'Enter') {
            handleSearch();
        }
    });

    // Enter key event for Language selection input
    $('#language-select').on('keydown', function (event) {
        if (event.key === 'Enter') {
            handleSearch();
        }
    });

    // Scroll to the top of #language-select when it loses focus
    $('#language-select').blur(function () {
        $('#language-select').scrollTop(0);
    });

    /* Function to handle readership search
    retrieve language and limit query
    go to the constructed href.
    If limit is empty don't include it */
    function handleReadershipSearch() {
        var language = $('#language-select-readership').val();
        var limit = $('#number-input-readership').val();
        if (limit != "") {
            window.location.href = '/readership/' + language + '/?limit=' + limit;
        } else {
            window.location.href = '/readership/' + language;
        }

    }

    // Click event for readership search button
    $('#search-button-readership').click(function () {
        handleReadershipSearch();
    });

    // Enter key event for readership number input
    $('#number-input-readership').on('keydown', function (event) {
        if (event.key === 'Enter') {
            handleReadershipSearch();
        }
    });

    /* Function to handle bookcount search
    retrieve language query
    go to the constructed href. */
    function handleBookcountSearch() {
        var language = $('#language-select-bc').val();
        window.location.href = '/bookcount?languages=' + language;
    }

    // Click event for bookcount search button
    $('#search-button-bc').click(function () {
        handleBookcountSearch();
    });

    // Enter key event for bookcount language select
    $('#language-select-bc').on('keydown', function (event) {
        if (event.key === 'Enter') {
            handleBookcountSearch();
        }
    });

    // Scroll to the top of #language-select-bc when it loses focus
    $('#language-select-bc').blur(function () {
        $('#language-select-bc').scrollTop(0);
    });

    //Function to handle transition primarily due to dynamic changes to size of elements
    function handleTransitionEffect() {
        const children = document.querySelectorAll('.book');

        // Trigger transition for each child
        children.forEach(child => {
            child.classList.add('move');
        });

        // Reset transition after animation ends
        setTimeout(() => {
            children.forEach(child => {
                child.classList.remove('move');
            });
        }, 500); // Adjust timeout to match transition duration in css
    }

    /* Click function on the open-close-menu class element
    On click the sidebar class is toggling the class active which
    triggers css animation.
    We also update the menu icon depeding on if the sidebar is active or not.
    We then handle transitions that are needed due to the expanding sidebar.
    */
    $('.open-close-menu').click(function () {
        $('.sidebar').toggleClass('active');
        if ($('.sidebar').hasClass('active')) {
            $('.open-close-menu').html('<i class="fa fa-times" aria-hidden="true"></i>');
        } else {
            $('.open-close-menu').html('<i class="fa fa-bars"></i>');
        }
        handleTransitionEffect();
    });

    // Function to categorize status codes
    function categorizeStatusCode(statusCode) {
        if (statusCode === "200") {
            return 'status-ok'; // Success status
        } else if (statusCode.startsWith("4")) {
            return 'status-warning'; // Client error status
        } else if (statusCode.startsWith("5")) {
            return 'status-error'; // Server error status
        } else {
            return 'status-info'; // Other status
        }
    }

    // Loop through each endpoint status span element
    $('.endpoint').each(function () {
        // Extract status code and status text from the text content
        var status = $(this).text().trim();
        var parts = status.split(' ');
        var statusCode = parts[0];

        // Get category based on status code
        var category = categorizeStatusCode(statusCode);

        // Add category class to the endpoint element
        $(this).addClass(category);
    });

    // Window resize handler to handle transitions when resizing
    window.addEventListener('resize', handleTransitionEffect);

    // Update the select option text based on selected languages
    $('#language-select').change(function () {
        var selectedLanguages = $(this).val();
        var selectLanguageOption = $('#select-language-option');
        if (selectedLanguages && selectedLanguages.length > 0) {
            var languageNames = selectedLanguages.map(function (lang) {
                return $('#language-select option[value="' + lang + '"]').val();
            }).join(',');
            selectLanguageOption.text(languageNames);
        } else {
            selectLanguageOption.text("Select Languages");
        }
    });

    $('#language-select-bc').change(function () {
        var selectedLanguages = $(this).val();
        var selectLanguageOption = $('#select-language-option');
        if (selectedLanguages && selectedLanguages.length > 0) {
            var languageNames = selectedLanguages.map(function (lang) {
                return $('#language-select-bc option[value="' + lang + '"]').val();
            }).join(',');
            selectLanguageOption.text(languageNames);
        } else {
            selectLanguageOption.text("Select Languages");
        }
    });

    // Scroll to top of language select when page loads
    $('#language-select').scrollTop(0);
    $('#language-select-bc').scrollTop(0);
});