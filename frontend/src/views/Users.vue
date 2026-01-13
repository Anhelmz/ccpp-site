<template>
  <div class="px-4 sm:px-0">
    <div class="max-w-6xl mx-auto">
      <!-- Header -->
      <div class="flex justify-between items-center mb-6">
        <h1 class="text-3xl font-bold text-gray-900">Users Management</h1>
        <button class="btn-primary" @click="showCreateModal = true">
          Add New User
        </button>
      </div>

      <!-- Error Message -->
      <div
        v-if="error"
        class="mb-4 bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded"
      >
        {{ error }}
      </div>

      <!-- Loading State -->
      <div v-if="loading" class="text-center py-8">
        <div
          class="inline-block animate-spin rounded-full h-8 w-8 border-b-2 border-main"
        ></div>
        <p class="mt-2 text-gray-600">Loading users...</p>
      </div>

      <!-- Users Table -->
      <div v-else class="card">
        <div v-if="users.length === 0" class="text-center py-8">
          <p class="text-gray-500 text-lg">
            No users found. Create your first user!
          </p>
        </div>

        <div v-else class="overflow-x-auto">
          <table class="min-w-full divide-y divide-gray-200">
            <thead class="bg-gray-50">
              <tr>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  ID
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Name
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Email
                </th>
                <th
                  class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Created At
                </th>
                <th
                  class="px-6 py-3 text-right text-xs font-medium text-gray-500 uppercase tracking-wider"
                >
                  Actions
                </th>
              </tr>
            </thead>
            <tbody class="bg-white divide-y divide-gray-200">
              <tr v-for="user in users" :key="user.id" class="hover:bg-gray-50">
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">
                  {{ user.id }}
                </td>
                <td
                  class="px-6 py-4 whitespace-nowrap text-sm font-medium text-gray-900"
                >
                  {{ user.name }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ user.email }}
                </td>
                <td class="px-6 py-4 whitespace-nowrap text-sm text-gray-500">
                  {{ formatDate(user.created_at) }}
                </td>
                <td
                  class="px-6 py-4 whitespace-nowrap text-right text-sm font-medium space-x-2"
                >
                  <button
                    class="text-main hover:opacity-80"
                    @click="editUser(user)"
                  >
                    Edit
                  </button>
                  <button
                    class="text-red-600 hover:text-red-900"
                    @click="deleteUser(user.id)"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>

    <!-- Create/Edit User Modal -->
    <div
      v-if="showCreateModal || showEditModal"
      class="fixed inset-0 bg-gray-600 bg-opacity-50 overflow-y-auto h-full w-full z-50"
    >
      <div
        class="relative top-20 mx-auto p-5 border w-96 shadow-lg rounded-md bg-white"
      >
        <div class="mt-3">
          <h3 class="text-lg font-medium text-gray-900 mb-4">
            {{ showCreateModal ? "Create New User" : "Edit User" }}
          </h3>

          <form @submit.prevent="submitForm">
            <div class="mb-4">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Name
              </label>
              <input
                v-model="form.name"
                type="text"
                required
                class="input-field"
                placeholder="Enter user name"
              />
            </div>

            <div class="mb-6">
              <label class="block text-sm font-medium text-gray-700 mb-2">
                Email
              </label>
              <input
                v-model="form.email"
                type="email"
                required
                class="input-field"
                placeholder="Enter user email"
              />
            </div>

            <div class="flex justify-end space-x-3">
              <button type="button" class="btn-secondary" @click="closeModal">
                Cancel
              </button>
              <button
                type="submit"
                :disabled="submitting"
                class="btn-primary"
                :class="{ 'opacity-50 cursor-not-allowed': submitting }"
              >
                {{
                  submitting
                    ? "Saving..."
                    : showCreateModal
                      ? "Create"
                      : "Update"
                }}
              </button>
            </div>
          </form>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { ref, onMounted } from "vue";
import { userService } from "@/services/userService";

export default {
  name: "Users",
  setup() {
    const users = ref([]);
    const loading = ref(true);
    const error = ref("");
    const showCreateModal = ref(false);
    const showEditModal = ref(false);
    const submitting = ref(false);
    const editingUser = ref(null);

    const form = ref({
      name: "",
      email: "",
    });

    const fetchUsers = async () => {
      try {
        loading.value = true;
        error.value = "";
        users.value = await userService.getUsers();
      } catch (err) {
        error.value = err.message;
      } finally {
        loading.value = false;
      }
    };

    const submitForm = async () => {
      try {
        submitting.value = true;
        error.value = "";

        if (showCreateModal.value) {
          await userService.createUser({
            name: form.value.name,
            email: form.value.email,
          });
        } else {
          await userService.updateUser(editingUser.value.id, {
            name: form.value.name,
            email: form.value.email,
          });
        }

        await fetchUsers();
        closeModal();
      } catch (err) {
        error.value = err.message;
      } finally {
        submitting.value = false;
      }
    };

    const editUser = (user) => {
      editingUser.value = user;
      form.value = {
        name: user.name,
        email: user.email,
      };
      showEditModal.value = true;
    };

    const deleteUser = async (id) => {
      if (confirm("Are you sure you want to delete this user?")) {
        try {
          error.value = "";
          await userService.deleteUser(id);
          await fetchUsers();
        } catch (err) {
          error.value = err.message;
        }
      }
    };

    const closeModal = () => {
      showCreateModal.value = false;
      showEditModal.value = false;
      editingUser.value = null;
      form.value = {
        name: "",
        email: "",
      };
    };

    const formatDate = (dateString) => {
      return new Date(dateString).toLocaleDateString("en-US", {
        year: "numeric",
        month: "short",
        day: "numeric",
        hour: "2-digit",
        minute: "2-digit",
      });
    };

    onMounted(() => {
      fetchUsers();
    });

    return {
      users,
      loading,
      error,
      showCreateModal,
      showEditModal,
      submitting,
      form,
      fetchUsers,
      submitForm,
      editUser,
      deleteUser,
      closeModal,
      formatDate,
    };
  },
};
</script>
