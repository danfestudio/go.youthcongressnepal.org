document.addEventListener("DOMContentLoaded", function() {
    fetch('/static/js/location.json')  // Ensure the correct path
        .then(response => response.json())
        .then(data => {
            const districtsMap = {};  // Map to hold districts, palikas, and wards

            // Process the data to structure it by district, palika, and ward
            data.forEach(location => {
                if (!districtsMap[location.district]) {
                    districtsMap[location.district] = {};
                }

                if (!districtsMap[location.district][location.palika]) {
                    districtsMap[location.district][location.palika] = {
                        type: location.type,
                        wards: []
                    };
                }

                districtsMap[location.district][location.palika].wards.push(location.ward);
            });

            // Populate the District dropdown
            const pDistrictSelect = document.getElementById("p_district");
            const tDistrictSelect = document.getElementById("t_district");

            Object.keys(districtsMap).forEach(district => {
                const option = document.createElement("option");
                option.value = district;
                option.textContent = district;
                pDistrictSelect.appendChild(option);
                tDistrictSelect.appendChild(option.cloneNode(true));
            });

            pDistrictSelect.addEventListener("change", function() {
                updatePalikaDropdown(this.value, "p_palika", districtsMap);
                updateWardDropdown("p_wada", []);
            });

            tDistrictSelect.addEventListener("change", function() {
                updatePalikaDropdown(this.value, "t_palika", districtsMap);
                updateWardDropdown("t_wada", []);
            });

            document.getElementById("p_palika").addEventListener("change", function() {
                updateWardDropdown("p_wada", districtsMap[pDistrictSelect.value][this.value].wards);
            });

            document.getElementById("t_palika").addEventListener("change", function() {
                updateWardDropdown("t_wada", districtsMap[tDistrictSelect.value][this.value].wards);
            });

            function updatePalikaDropdown(district, palikaSelectId, districtsMap) {
                const palikaSelect = document.getElementById(palikaSelectId);
                palikaSelect.innerHTML = '<option value="" selected>Select your palika</option>';

                if (districtsMap[district]) {
                    Object.keys(districtsMap[district]).forEach(palika => {
                        const option = document.createElement("option");
                        option.value = palika;
                        option.textContent = `${palika} (${districtsMap[district][palika].type})`;
                        palikaSelect.appendChild(option);
                    });
                }
            }

            function updateWardDropdown(wardSelectId, wards) {
                const wardSelect = document.getElementById(wardSelectId);
                wardSelect.innerHTML = '<option value="" selected>Select your ward</option>';

                wards.forEach(ward => {
                    const option = document.createElement("option");
                    option.value = ward;
                    option.textContent = `Ward ${ward}`;
                    wardSelect.appendChild(option);
                });
            }
        })
        .catch(error => console.error("Error fetching location data:", error));
});

document.addEventListener("DOMContentLoaded", function() {
    const password = document.getElementById("password");
    const confirmPassword = document.getElementById("confirm_password");
    const passwordError = document.getElementById("password_error");

    // Function to check if passwords match
    function checkPasswordsMatch() {
        // Show error only if both fields have values and passwords do not match
        if (password.value !== confirmPassword.value && confirmPassword.value !== "") {
            passwordError.style.display = "inline";  // Show error if passwords don't match
        } else {
            passwordError.style.display = "none";  // Hide error if passwords match
        }
    }

    // Add event listeners to check password match on input change
    password.addEventListener("input", checkPasswordsMatch);
    confirmPassword.addEventListener("input", checkPasswordsMatch);

    // Optional: You can prevent form submission if passwords don't match
    const form = document.getElementById("registration_form");
    form.addEventListener("submit", function(event) {
        if (password.value !== confirmPassword.value) {
            event.preventDefault(); // Prevent form submission
            passwordError.style.display = "inline";  // Show error if passwords don't match
        }
    });
});

